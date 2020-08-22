pub mod collector;

/// Run a command within a shell with a specified directory
pub fn run_command(full_cmd: &str, dir: &str) -> Result<(String, String), String> {
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

pub fn run_command_in_background(full_cmd: &str, dir: &str) {
    use std::process::Command;
    let mut splot = full_cmd.split(" ").collect::<Vec<&str>>();
    splot.reverse();
    let cmd = splot.pop().expect(&format!(
        "Something went wrong when splitting command: \"{}\"",
        full_cmd
    ));
    splot.reverse();
    Command::new(cmd)
        .current_dir(dir)
        .args(splot)
        .spawn()
        .unwrap();
}
