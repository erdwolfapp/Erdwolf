use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RegistrationTemplate {
    pub invite_code_required: bool,
    pub invite_code: String,
    pub error_invite_code_not_specified: bool,
    pub error_invalid_invite_code: bool,
    pub error_overused_invite_code: bool,
    pub error_user_already_exists: bool,
    pub error_confirmation_pass_doesnt_match: bool,
}

impl RegistrationTemplate {
    pub fn new() -> RegistrationTemplate {
        RegistrationTemplate {
            invite_code_required: false,
            invite_code: "".to_string(),
            error_invite_code_not_specified: false,
            error_invalid_invite_code: false,
            error_overused_invite_code: false,
            error_user_already_exists: false,
            error_confirmation_pass_doesnt_match: false,
        }
    }
}
