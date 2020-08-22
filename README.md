# Erdwolf
This is a web-service that allows remote management of podman containers over a web-interface, similar to what a `Dokku` does, just more accessible.

The project is based on `Rust` and uses [rocket.rs](https://rocket.rs) for the webserver and [diesel.rs](https://diesel.rs) for database ORM.

# Building
You can build the server from sources for yourself.

## Prerequisites
- Rust nightly (1.47+) (Older versions might work, but weren't tested)
- SQLite libraries (`sqlite-sys` needs them)

## Installing rust
1.    Follow steps on [rustup.rs](https://rustup.rs)
2.    Set `nightly` toolchain as default

      ```rustup default nightly```
3.    You're all set up

## Building process

- Clone this repository
    ```
    git clone https://github.com/erdwolfapp/Erdwolf.git
    ```

- Install `diesel-cli` which you're going to need for DB creation. Feel free to remove arguments if you want to use it for other types of databases.
    ```
    cargo install diesel-cli --no-default-features --features="sqlite"
    ```

- Add `~/.cargo/bin` to your PATH (If you're on Linux)
- Run diesel migrations to create `erdwolf.db`
    ```
    diesel migration run
    ```
- Build the application
    ```
    cargo build --release
    ```

# Running Erdwolf
## Running from precompiled
TBD
## Running from sources
- Clone this repository
    ```
    git clone https://github.com/erdwolfapp/Erdwolf.git
    ```

- Install `diesel-cli` which you're going to need for DB creation
    ```
    cargo install diesel-cli
    ```

- Add `~/.cargo/bin` to your PATH (If you're on Linux)
- Run diesel migrations to create `erdwolf.db`
    ```
    diesel migration run
    ```
- Build the application
    ```
    cargo run
    ```