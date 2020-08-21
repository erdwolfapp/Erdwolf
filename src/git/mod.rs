struct GitServer {}

#[allow(unused)]
impl GitServer {
    pub fn new() -> GitServer {
        GitServer {}
    }

    /// Returns the version of GIT currently installed.
    /// For example: "2.28.0"
    pub fn get_local_git_version(&self) -> String {
        let msg = self.run_command("git --version", "/");
        if let Ok(msg) = msg {
            msg.0.split("version ").collect::<Vec<&str>>()[1]
                .trim()
                .to_string()
        } else {
            println!("ERROR: No Git available locally!");
            "None".to_string()
        }
    }

    /// Creates a new bare git repository with a specified name
    /// Repository home is taken from environment variable "ERDWOLF_REPO_HOME"
    /// Hostname is from "ERDWOLF_HOSTNAME"
    pub fn create_new_repository(&self, project: &str) -> Result<String, String> {
        let repo_dir = std::path::Path::new(dotenv!("ERDWOLF_REPO_HOME"));
        let hostname = dotenv!("ERDWOLF_HOSTNAME");
        let project_path = format!("{}/{}", repo_dir.display(), project);
        let project_path = std::path::Path::new(&project_path);
        let project_dir = std::fs::File::open(project_path);
        if let Ok(_) = project_dir {
            return Err("Project with that name already exists".to_string());
        } else {
            std::fs::create_dir_all(project_path).expect(&format!(
                "Unable to create a folder at: \"{}\"\nVerify that you have permissions!",
                project_path.display()
            ));
            let msg = self.run_command("git init --bare", &project_path.display().to_string());
            if let Err(msg) = msg {
                return Err(format!(
                    "Something went wrong during project initialization: \n{}",
                    msg
                ));
            }
            Ok(format!("{}:{}", hostname, project_path.display()))
        }
    }

    /// Run a command within a shell with a specified directory
    fn run_command(&self, full_cmd: &str, dir: &str) -> Result<(String, String), String> {
        use std::process::Command;
        let mut splot = full_cmd.split(" ").collect::<Vec<&str>>();
        splot.reverse();
        let cmd = splot.pop().expect(&format!(
            "Something went wrong when splitting command: \"{}\"",
            full_cmd
        ));
        splot.reverse();
        match Command::new(cmd).current_dir(dir).args(splot).output() {
            Ok(output) => Ok((
                String::from_utf8_lossy(&output.stdout).to_string(),
                String::from_utf8_lossy(&output.stderr).to_string(),
            )),
            Err(_) => Err(format!(
                "Something went wrong running command: \"{}\"!",
                full_cmd
            )),
        }
    }
}

#[test]
fn test_run_command() {}

#[test]
fn test_get_local_git_version() {
    let server = GitServer::new();
    assert_ne!(server.get_local_git_version(), "None".to_string());
    assert_eq!(server.get_local_git_version(), "2.28.0".to_string());
}

#[test]
fn test_create_new_repository() {
    let server = GitServer::new();
    let out = server.create_new_repository("test_project");
    assert!(out.is_ok());
    let repo_dir = std::path::Path::new(dotenv!("ERDWOLF_REPO_HOME"));
    std::fs::remove_dir_all(format!("{}/{}", repo_dir.display(), "test_project")).unwrap();
    assert!(true);
}
