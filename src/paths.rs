use crate::db::{
    establish_connection,
    models::{InviteCode, User, UserInsertable},
};
use crate::diesel::ExpressionMethods;
use crate::forms::{LoginForm, RegisterForm};
use crate::template_models::RegistrationTemplate;
use diesel::{QueryDsl, RunQueryDsl};
use dotenv::dotenv;
use passwors::*;
use rocket::http::Cookies;
use rocket::request::Form;
use rocket::response::Redirect;
use rocket_contrib::templates::Template;
use std::env;

#[get("/")]
pub fn index() -> Template {
    Template::render("index", 0u32)
}

#[get("/login")]
pub fn login() -> Template {
    Template::render("login", 0u32)
}

#[get("/register/<code>")]
pub fn register(code: Option<String>) -> Result<Template, Redirect> {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED") == "1" {
        let mut registration = RegistrationTemplate {
            invite_code_required: false,
            invite_code: "".to_string(),
        };
        if env::var("REGISTRATION_INVITE_CODE_ONLY") == "1" {
            registration.invite_code_required = true;
            if let Some(code) = code {
                registration.invite_code = code;
            }
        }
        return Template::render("register", registration);
    }
    Redirect::to(uri!(index))
}

#[get("/home")]
pub fn home(mut cookies: Cookies) -> Result<Template, Redirect> {
    if let Some(sid) = cookies.get_private("SID") {
        Ok(Template::render("home", 0u32))
    } else {
        Err(Redirect::to(uri!(index)))
    }
}

#[post("/api/login", data = "<form>")]
pub fn login_api(mut cookies: Cookies, form: Form<LoginForm>) -> Redirect {
    let connection = establish_connection();
    use crate::schema::users::dsl::*;
    let results = users
        .filter(username.eq(form.username.to_string()))
        .load::<User>(&connection)
        .expect("Error loading users");
    if results.len() == 1 {
        let user = results.get(0).unwrap();
        let pw = ClearTextPassword::from_string(form.password.to_string()).unwrap();
        let salt = user.password_salt;
        let a2hash = Argon2PasswordHasher::default();
        let pw_hash = pw.hash(&a2hash, &salt);
        if user.password_hash == pw_hash {
            Redirect::to(uri!(home))
        } else {
            Redirect::to(uri!(login))
        }
    } else {
        Redirect::to(uri!(login))
    }
}

#[post("/api/register", data = "<form>")]
pub fn register_api(mut cookies: Cookies, form: Form<RegisterForm>) -> Redirect {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED") == "1" {
        let invite_code = None;
        if env::var("REGISTRATION_INVITE_CODE_ONLY") == "1" {
            if form.invite_code.len() == 0 {
                //TODO: Invite code wasn't specified
                return Redirect::to(uri!(register));
            }
            use crate::schema::invite_codes::dsl::*;
            let results = invite_codes
                .filter(code.eq(form.invite_code.to_string()))
                .load::<InviteCode>(&connection)
                .expect("Error loading invite codes");
            if results.len() != 1 {
                //TODO: Invalid invite_code
                return Redirect::to(uri!(register));
            }
            let result = results.get(0).unwrap();
            if result.max_usages <= result.times_used {
                //TODO: Overused invite_code
                return Redirect::to(uri!(register));
            }
            diesel::update(invite_codes.filter(code.eq(form.invite_code)))
                .set(times_used.eq(result.times_used + 1))
                .expect(&format!("Unable to find invite_code {}", id));
            invite_code = Some(form.invite_code);
        }
        let connection = crate::db::establish_connection();
        use crate::schema::users::dsl::*;
        let results = users
            .filter(username.eq(form.username.to_string()))
            .load::<User>(&connection)
            .expect("Error loading users");
        if results.len() >= 1 || form.password != form.repeat_password {
            //TODO: Wrong confirmation password or user already exists
            return Redirect::to(uri!(register));
        } else {
            let pw = ClearTextPassword::from_string(form.password.to_string()).unwrap();
            let salt = HashSalt::new().unwrap();
            let a2hash = Argon2PasswordHasher::default();
            let pw_hash = pw.hash(&a2hash, &salt);
            let new_user = UserInsertable {
                username: form.username,
                password_salt: salt,
                password_hash: pw_hash,
                invite_code_used: invite_code,
            };
            diesel::insert_into(users::table)
                .values(&new_user)
                .get_result(&connection)
                .expect("Something went wrong while inserting new user!");
            return Redirect::to(uri!(home));
        }
    } else {
        //TODO: Registration isn't enabled
        return Redirect::to(uri!(index));
    }
}
