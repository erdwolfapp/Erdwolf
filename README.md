# Erdwolf
This is a web-service that allows remote management of podman containers over a web-interface, similar to what a `Dokku` does, just more accessible.

The project is based on `Rust` and uses [rocket.rs](https://rocket.rs) for the webserver and [diesel.rs](https://diesel.rs) for database ORM.

# Building
You can build the server from sources for yourself.

## Prerequisites
- [Few basic requirements](#basic-requirements)
- Rust nightly (1.47+) (Older versions might work, but weren't tested)
- SQLite libraries (`sqlite-sys` needs them)
- OpenSSL

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

- Install `diesel_cli` which you're going to need for DB creation. Feel free to remove arguments if you want to use it for other types of databases.
    ```
    cargo install diesel_cli --no-default-features --features="sqlite"
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

- Install `diesel_cli` which you're going to need for DB creation
    ```
    cargo install diesel_cli --no-default-features --features="sqlite"
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


## Basic requirements
Requirements have been tested on bare docker containers for the matching distribution. All of them might not be neccessary for a successful build.

If you find that packages aren't needed, please post an issue.
### Arch linux
```
# pacman -S curl git clang pkg-config openssl sqlite
```

### Void Linux
```
# xbps-install curl git clang gcc pkg-config libclang libargon2-devel libressl-devel sqlite-devel
```

### Debian/Ubuntu
```
# apt install curl git clang libsqlite3-dev librust-openssl-dev
```