table! {
    containers (id) {
        id -> Integer,
        owner -> Integer,
        name -> Text,
        path -> Text,
        subdomain -> Nullable<Text>,
    }
}

table! {
    groups (id) {
        id -> Integer,
        group_name -> Text,
    }
}

table! {
    invite_codes (id) {
        id -> Integer,
        code -> Text,
        times_used -> Integer,
        max_usages -> Integer,
        generated -> Integer,
    }
}

table! {
    ports (id) {
        id -> Integer,
        port_number -> Integer,
        container -> Integer,
    }
}

table! {
    roles (id) {
        id -> Integer,
        name -> Nullable<Text>,
        permission_level -> Integer,
    }
}

table! {
    users (id) {
        id -> Integer,
        username -> Text,
        password_hash -> Text,
        password_salt -> Text,
        invite_code_used -> Nullable<Integer>,
        role -> Nullable<Integer>,
    }
}

table! {
    users_groups (id) {
        id -> Integer,
        user -> Integer,
        group -> Integer,
    }
}

joinable!(containers -> groups (owner));
joinable!(users -> roles (role));
joinable!(users_groups -> groups (group));
joinable!(users_groups -> users (user));

allow_tables_to_appear_in_same_query!(
    containers,
    groups,
    invite_codes,
    ports,
    roles,
    users,
    users_groups,
);
