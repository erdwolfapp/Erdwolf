package com.erdwolf.web

import io.ktor.locations.Location
import io.ktor.locations.get
import io.ktor.routing.Route

@Location("/")
class Root()

fun Route.root() {
    get<Root> {

    }
}