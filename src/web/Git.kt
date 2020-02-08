package com.erdwolf.web

import com.erdwolf.git.GitServer
import io.ktor.application.call
import io.ktor.http.HttpStatusCode
import io.ktor.response.respond
import io.ktor.routing.Route
import io.ktor.routing.get

lateinit var gitServer: GitServer

fun Route.git() {
    gitServer = GitServer()
    get("/git/new/{project}"){
        val project = call.parameters["project"]
        if(project == null) {
            call.respond(HttpStatusCode.BadRequest)
            return@get
        }
        if(!project.matches("[A-z0-9_-]+".toRegex())){
            call.respond(HttpStatusCode.BadRequest)
            return@get
        }
        val (err, msg) = gitServer.createNewRepository("$project.git")
        if(err) {
            call.respond(HttpStatusCode.InternalServerError, msg)
            return@get
        }
        call.respond(HttpStatusCode.OK, msg)
    }
    get("/git/{...}"){
        call.respond(HttpStatusCode.BadRequest)
    }
}