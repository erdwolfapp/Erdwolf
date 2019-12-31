package com.erdwolf.api.v1

import io.ktor.application.call
import io.ktor.locations.Location
import io.ktor.locations.get
import io.ktor.response.respondText
import io.ktor.routing.Route

@Location("/api/v1/check_health")
class HealthCheck()

fun Route.healthCheck() {
    get<HealthCheck> {
        // Check databases/other services.
        call.respondText("OK")
    }
}