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
pub fn index(mut cookies: Cookies) -> Redirect {
    if let Some(_sid) = cookies.get_private("UID") {
        Redirect::to(uri!(home))
    } else {
        Redirect::to(uri!(login))
    }
}

#[get("/login")]
pub fn login() -> Template {
    Template::render("login", 0u32)
}

#[get("/register?<error>")]
pub fn register(error: Option<u32>) -> Result<Template, Redirect> {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED").is_ok() {
        if env::var("REGISTRATION_ENABLED").unwrap() == "1" {
            let mut registration = RegistrationTemplate::new();
            if env::var("REGISTRATION_INVITE_CODE_ONLY").is_ok() {
                if env::var("REGISTRATION_INVITE_CODE_ONLY").unwrap() == "1" {
                    registration.invite_code_required = true;
                }
            }
            get_error(&mut registration, error);
            return Ok(Template::render("register", registration));
        }
    }
    Err(Redirect::to(uri!(index)))
}

fn get_error(template: &mut RegistrationTemplate, error_code: Option<u32>) {
    if let Some(error_code) = error_code {
        match error_code {
            0 => template.error_invite_code_not_specified = true,
            1 => template.error_invalid_invite_code = true,
            2 => template.error_overused_invite_code = true,
            3 => template.error_user_already_exists = true,
            4 => template.error_confirmation_pass_doesnt_match = true,
            _ => unimplemented!("Apparently we have a new error of number {}", error_code),
        }
    }
}

#[get("/register/<code>?<error>")]
pub fn register_with_code(code: Option<String>, error: Option<u32>) -> Result<Template, Redirect> {
    dotenv().ok();
    if env::var("REGISTRATION_ENABLED").is_ok() {
        if env::var("REGISTRATION_ENABLED").unwrap() == "1" {
            let mut registration = RegistrationTemplate::new();
            if env::var("REGISTRATION_INVITE_CODE_ONLY").is_ok() {
                if env::var("REGISTRATION_INVITE_CODE_ONLY").unwrap() == "1" {
                    registration.invite_code_required = true;
                    if let Some(code) = code {
                        registration.invite_code = code;
                    }
                }
            }
            get_error(&mut registration, error);
            return Ok(Template::render("register", registration));
        }
    }
    Err(Redirect::to(uri!(index)))
}

#[get("/home")]
pub fn home(mut cookies: Cookies) -> Result<Template, Redirect> {
    if let Some(_sid) = cookies.get_private("UID") {
        // Add home stuff
        Ok(Template::render("home", 0u32))
    } else {
        Err(Redirect::to(uri!(login)))
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
pub fn register_api(mut cookies: Cookies, form: Form<RegisterForm>) -> Redirect {
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
                if form.invite_code.is_none() || form.invite_code.unwrap().len() == 0 {
                    // Invite code wasn't specified
                    return Redirect::to(uri!(register: 0));
                }
                use crate::schema::invite_codes::dsl::*;
                let results = invite_codes
                    .filter(code.eq(form.invite_code.unwrap().to_string()))
                    .load::<InviteCode>(&connection)
                    .expect("Error loading invite codes");
                if results.len() != 1 {
                    // Invalid invite_code
                    let c = &form.invite_code.unwrap().to_string();
                    return Redirect::to(uri!(register_with_code: c, 1));
                }
                let result = results.get(0).unwrap();
                if result.max_usages <= result.times_used {
                    // Overused invite_code
                    return Redirect::to(uri!(register: 2));
                }
                diesel::update(invite_codes.filter(code.eq(form.invite_code.unwrap().to_string())))
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
        if results.len() >= 1 {
            // User already exists
            return Redirect::to(uri!(register: 3));
        }
        if form.password != form.repeat_password {
            // Confirmation password doesn't match
            return Redirect::to(uri!(register: 4));
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
            let results = users
                .filter(username.eq(form.username.to_string()))
                .load::<User>(&connection)
                .expect("Error loading users");
            let result = results.get(0).unwrap();
            if results.len() == 1 {
                cookies.add_private(Cookie::new("UID", result.id.to_string()));
            }
            return Redirect::to(uri!(home));
        }
    } else {
        // Registration isn't enabled
        return Redirect::to(uri!(index));
    }
}
