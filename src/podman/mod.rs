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
    fn start_server(&self) -> Result<(), String> {
        if std::fs::File::open(format!("/run/user/{}/podman/podman.sock", self.uid)).is_err() {
            run_command_in_background("podman system service -t 5000", "/");
            sleep(Duration::from_millis(1));
        }
        Ok(())
    }

    pub fn get_server_info(&self) {
        self.start_server();
        let out = self.send_request_to_podman("http://d/v1.0.0/libpod/info");
        if let Ok(s) = out {
            println!("Sock response: \"{}\"", s);
        } else {
            println!("Error `{}` occured!", out.unwrap_err());
        }
    }

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

    fn send_request_to_podman(&self, url: &str) -> Result<String, u32> {
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

#[test]
fn get_server_info_test() {
    let podman = Podman::new();
    podman.get_server_info();
}

#[test]
fn get_uid_test() {
    let mut podman = Podman { uid: 9999 };
    let id = podman.get_uid();
    assert!(id.is_ok());
    assert_ne!(id.unwrap(), 9999);
}
