use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RegistrationTemplate {
    pub invite_code_required: bool,
    pub invite_code: String,
}
