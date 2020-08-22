#![feature(proc_macro_hygiene, decl_macro)]
#[macro_use]
extern crate rocket;
#[macro_use]
extern crate dotenv_codegen;
#[macro_use]
extern crate diesel;

pub mod db;
pub mod forms;
pub mod git;
pub mod paths;
pub mod podman;
pub mod schema;
pub mod template_models;
pub mod utils;

use rocket_contrib::serve::StaticFiles;
use rocket_contrib::templates::Template;

fn main() {
    rocket::ignite()
        .attach(Template::fairing())
        .mount("/public", StaticFiles::from("css"))
        .mount(
            "/",
            routes![
                paths::index,
                paths::login,
                paths::register,
                paths::register_with_code,
                paths::home,
                paths::login_api,
                paths::register_api
            ],
        )
        .launch();
}
