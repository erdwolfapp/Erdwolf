use crate::db::{
    establish_connection,
    models::{InviteCode, User, UserInsertable},
};
use crate::diesel::ExpressionMethods;
use crate::forms::{LoginForm, RegisterForm};
use crate::template_models::RegistrationTemplate;
use argonautica::{Hasher, Verifier};
use diesel::{QueryDsl, RunQueryDsl};
use dotenv::dotenv;
use rand::prelude::*;
use rocket::http::{Cookie, Cookies};
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

#[get("/register")]
pub fn register() -> Result<Template, Redirect> {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED").is_ok() {
        if env::var("REGISTRATION_ENABLED").unwrap() == "1" {
            let mut registration = RegistrationTemplate {
                invite_code_required: false,
                invite_code: "".to_string(),
            };
            if env::var("REGISTRATION_INVITE_CODE_ONLY").is_ok() {
                if env::var("REGISTRATION_INVITE_CODE_ONLY").unwrap() == "1" {
                    registration.invite_code_required = true;
                }
            }
            return Ok(Template::render("register", registration));
        }
    }
    Err(Redirect::to(uri!(index)))
}

#[get("/register/<code>")]
pub fn register_with_code(code: Option<String>) -> Result<Template, Redirect> {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED").is_ok() {
        if env::var("REGISTRATION_ENABLED").unwrap() == "1" {
            let mut registration = RegistrationTemplate {
                invite_code_required: false,
                invite_code: "".to_string(),
            };
            if env::var("REGISTRATION_INVITE_CODE_ONLY").is_ok() {
                if env::var("REGISTRATION_INVITE_CODE_ONLY").unwrap() == "1" {
                    registration.invite_code_required = true;
                    if let Some(code) = code {
                        registration.invite_code = code;
                    }
                }
            }
            return Ok(Template::render("register", registration));
        }
    }
    Err(Redirect::to(uri!(index)))
}

#[get("/home")]
pub fn home(mut cookies: Cookies) -> Result<Template, Redirect> {
    if let Some(_sid) = cookies.get_private("UID") {
        //TODO: Add home stuff
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
        let mut verifier = Verifier::default();
        let is_valid = verifier
            .with_hash(&user.password_hash)
            .with_password(form.password.to_string())
            .with_secret_key(&user.password_salt)
            .verify()
            .unwrap();
        if is_valid {
            cookies.add_private(Cookie::new("UID", user.id.to_string()));
            Redirect::to(uri!(home))
        } else {
            Redirect::to(uri!(login))
        }
    } else {
        Redirect::to(uri!(login))
    }
}

#[post("/api/register", data = "<form>")]
pub fn register_api(mut _cookies: Cookies, form: Form<RegisterForm>) -> Redirect {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED").is_err() {
        println!("Please make sure to set the REGISTRATION_ENABLED env variable!");
        std::process::exit(1);
    }
    let connection = crate::db::establish_connection();
    if env::var("REGISTRATION_ENABLED").unwrap() == "1" {
        let mut invite_code: Option<i32> = None;
        if env::var("REGISTRATION_INVITE_CODE_ONLY").is_ok() {
            if env::var("REGISTRATION_INVITE_CODE_ONLY").unwrap() == "1" {
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
                    let c = &form.invite_code.to_string();
                    return Redirect::to(uri!(register_with_code: c));
                }
                let result = results.get(0).unwrap();
                if result.max_usages <= result.times_used {
                    //TODO: Overused invite_code
                    return Redirect::to(uri!(register));
                }
                diesel::update(invite_codes.filter(code.eq(form.invite_code.to_string())))
                    .set(times_used.eq(result.times_used + 1))
                    .execute(&connection)
                    .expect(&format!("Unable to find post {}", result.id));
                invite_code = Some(result.id);
            }
        }
        use crate::schema::users::dsl::*;
        let results = users
            .filter(username.eq(form.username.to_string()))
            .load::<User>(&connection)
            .expect("Error loading users");
        if results.len() >= 1 || form.password != form.repeat_password {
            //TODO: Wrong confirmation password or user already exists
            return Redirect::to(uri!(register));
        } else {
            let salt: String = thread_rng()
                .sample_iter(&rand::distributions::Alphanumeric)
                .take(32)
                .collect();
            let mut hasher = Hasher::default();
            let hash = hasher
                .with_password(form.password.to_string())
                .with_secret_key(&salt)
                .hash()
                .unwrap();
            let new_user = UserInsertable {
                username: form.username,
                password_hash: &hash,
                password_salt: &salt,
                invite_code_used: invite_code,
                role: None,
            };
            diesel::insert_into(users)
                .values(&new_user)
                .execute(&connection)
                .expect("Something went wrong while inserting new user!");
            return Redirect::to(uri!(home));
        }
    } else {
        //TODO: Registration isn't enabled
        return Redirect::to(uri!(index));
    }
}
