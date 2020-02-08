package com.erdwolf.git

import io.ktor.locations.Location
import java.io.BufferedReader
import java.io.File
import java.io.InputStreamReader
import java.lang.RuntimeException


/*
    https://git-scm.com/book/en/v2/Git-on-the-Server-Setting-Up-the-Server
    Server is created using the local Git executable.
 */

//TODO: Change this into a commandline argument or config option
val repoDir = File(System.getProperty("user.home"), ".repos")
val hostname = "localhost"

class GitServer{
    init{
        println("Detected GIT version: "+getLocalGitVersion())
    }

    /*
    Return
    String: Git version

    Git version example: "2.24.1"
    */
    private fun getLocalGitVersion(): String{
        val (error, msg) = runCommand("git --version", null)
        if(error){
            throw RuntimeException("No Git available locally!")
        }else{
            return msg.split("version ")[1]
        }
    }

    /*
    Parameters
    @project: project to be created

    Return
    Boolean: Error
    String: Message (Error if boolean true, Command output if false)
    */
    // TODO: It'd be better if we implemented per-user or per-company repositories
    fun createNewRepository(project: String): Pair<Boolean, String>{
        val projectDir = File(repoDir, project)
        if(projectDir.exists())
            return true to "Project with that name already exists"
        projectDir.mkdirs()
        val (err, msg) = runCommand("git init --bare", projectDir)
//        val (err, msg) = runCommand("git init", projectDir)
        if(err)
            return true to "Something went wrong during project initialization: \n$msg"
        return false to "$hostname:"+projectDir.absolutePath
    }

    /*
        Parameters
        @cmd: Command to be executed
        @dir: Directory, in which the command executes

        Return
        Boolean: Error
        String: Message (Error if boolean true, Command output if false)
     */
    private fun runCommand(cmd: String, dir: File?): Pair<Boolean, String>{
        val p = if(dir != null) {
            println(dir.absolutePath+" exists: "+dir.exists())
            ProcessBuilder().command(cmd.split(" ")).directory(dir).start()
        }else{
//            ProcessBuilder().command(cmd).start()
            Runtime.getRuntime().exec(cmd)
        }
        val stdInput = BufferedReader(InputStreamReader(p.inputStream))
        val stdError = BufferedReader(InputStreamReader(p.errorStream))
        return if(stdError.readText().isNotEmpty()){
            true to stdError.readText()
        }else {
            false to stdInput.readText()
        }
    }
}