use crate::utils::{collector::Collector, run_command, run_command_in_background};
use curl::easy::Easy2;
use regex::Regex;
use std::thread::sleep;
use std::time::Duration;

struct Podman {
    uid: u32,
}

#[allow(unused)]
impl Podman {
    pub fn new() -> Podman {
        let mut podman = Podman { uid: 9999 };
        podman.get_uid();
        podman
    }

    // Starts the server up for 5000 seconds.
    /// Start the Podman server service for interaction
    fn start_server(&self) -> Result<(), String> {
        if std::fs::File::open(format!("/run/user/{}/podman/podman.sock", self.uid)).is_err() {
            run_command_in_background("podman system service -t 5000", "/");
            sleep(Duration::from_millis(1));
        }
        Ok(())
    }

    // System commands

    /// Get information about the Podman host
    pub fn get_server_info(&self) {
        self.start_server();
        let out = self.send_request_to_podman("http://d/v1.0.0/libpod/info", CurlRequestType::GET);
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
        } else {
            println!("Error `{}` occured!", out.unwrap_err());
        }
    }

    // Image commands

    /// Get list of all images
    pub fn get_images(&self) -> Result<String, String> {
        self.start_server();
        let out =
            self.send_request_to_podman("http://d/v1.0.0/libpod/images/json", CurlRequestType::GET);
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    // Container commands

    pub fn start_container(&self, name: String) -> Result<String, String> {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}/start", name),
            CurlRequestType::POST,
        );
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    pub fn restart_container(&self, name: String) -> Result<String, String> {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}/restart", name),
            CurlRequestType::POST,
        );
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    pub fn stop_container(&self, name: String) -> Result<String, String> {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}/stop", name),
            CurlRequestType::POST,
        );
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    pub fn inspect_container(&self, name: String) -> Result<String, String> {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}/json", name),
            CurlRequestType::GET,
        );
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    pub fn export_container(&self, name: String) -> Result<String, String> {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}/export", name),
            CurlRequestType::GET,
        );
        if let Ok(s) = out {
            // TODO: Save tarball as a file and don't print it to STDOUT
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    pub fn container_exists(&self, name: String) -> bool {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}/exists", name),
            CurlRequestType::GET,
        );
        if let Ok(_) = out {
            true
        } else {
            false
        }
    }

    pub fn delete_container(&self, name: String) -> Result<String, String> {
        self.start_server();
        let out = self.send_request_to_podman(
            &format!("http://d/v1.0.0/libpod/containers/{}", name),
            CurlRequestType::DELETE,
        );
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
            Ok(s)
        } else {
            Err(format!("Error `{}` occured!", out.unwrap_err()))
        }
    }

    // Utilities

    fn get_uid(&mut self) -> Result<u32, String> {
        let msg = run_command("id", "/");
        if let Err(msg) = msg {
            return Err(format!(
                "Something went wrong while acquiring userid: \n{}",
                msg
            ));
        }
        let msg = msg.unwrap().0;
        let re = Regex::new(r"^uid=([0-9]+)").unwrap();
        let recap = re.captures(&msg);
        if let Some(recap) = recap {
            self.uid = recap[1].parse().unwrap()
        } else {
            return Err("Something went wrong collecting UID.".to_string());
        }
        Ok(self.uid)
    }

    fn send_request_to_podman(&self, url: &str, reqtype: CurlRequestType) -> Result<String, u32> {
        let mut easy = Easy2::new(Collector(Vec::new()));
        easy.unix_socket(&format!("/run/user/{}/podman/podman.sock", self.uid))
            .unwrap();
        easy.url(url).unwrap();
        easy.perform().unwrap();
        if easy.response_code().unwrap() == 200 {
            let contents = easy.get_ref();
            Ok(String::from(String::from_utf8_lossy(&contents.0)))
        } else {
            Err(easy.response_code().unwrap())
        }
    }
}

enum CurlRequestType {
    GET,
    POST,
    DELETE,
}

#[test]
fn get_server_info_test() {
    let podman = Podman::new();
    podman.get_server_info();
}

#[test]
fn get_images_test() {
    let podman = Podman::new();
    assert!(podman.get_images().is_ok());
}

#[test]
fn get_uid_test() {
    let mut podman = Podman { uid: 9999 };
    let id = podman.get_uid();
    assert!(id.is_ok());
    assert_ne!(id.unwrap(), 9999);
}
