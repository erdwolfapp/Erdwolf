#[derive(Queryable)]
pub struct User {
    pub id: i32,
    pub username: String,
    pub password_hash: String,
    pub password_salt: String,
    pub invite_code_used: Option<i32>,
}

use crate::schema::users;
#[derive(Insertable)]
#[table_name = "users"]
pub struct UserInsertable<'a> {
    pub username: &'a str,
    pub password_hash: &'a str,
    pub password_salt: &'a str,
    pub invite_code_used: Option<i32>,
}

#[derive(Queryable)]
pub struct InviteCode {
    pub id: i32,
    pub code: String,
    pub times_used: i32,
    pub max_usages: i32,
    pub generated: i32,
}

use crate::schema::invite_codes;
#[derive(Insertable)]
#[table_name = "invite_codes"]
pub struct InviteCodeInsertable<'a> {
    pub code: &'a str,
    pub times_used: i32,
    pub max_usages: i32,
    pub generated: i32,
}
