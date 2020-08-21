use rocket::http::RawStr;

#[allow(unused)]
#[derive(FromForm)]
pub struct LoginForm<'f> {
    pub username: &'f RawStr,
    pub password: &'f RawStr,
}
#[allow(unused)]
#[derive(FromForm)]
pub struct RegisterForm<'f> {
    pub username: &'f RawStr,
    pub password: &'f RawStr,
    pub repeat_password: &'f RawStr,
    pub invite_code: Option<&'f RawStr>,
}
