# microwars
Microservices Wars

Blah blah intro - will get to this later.

# Specs
## Project structure
Each service will have its own folder, with the folder name being the character the developer chooses. A project README.md file should be present in the service's folder, to describe the stack used, along with possible startup parameters and setup info.

## Endpoints
Each of the attack endpoints (jab, cross, hook, uppercut) will also initiate requests to the relevant attack endpoints of the opponents, based on the combos section below.
### /status
Just returns "OK" and HTTP status code 200.
### /test
Will initiate a GET to the /status endpoint of the opponent - just to verify connectivity of the two opponents.
### /combat
Will start the combat, by sending a jab and hook request to the opponent.

### /jab
Returns a newly generated UUID along with the result of FIB(2)
### /cross
Returns a newly generated UUID along with the result of FIB(4)
### /hook
Returns a newly generated UUID along with the result of FIB(8)
### /uppercut
Returns a newly generated UUID along with the result of FIB(16)

## Combos
### jab
Respond with jab, jab
### cross
Respond with jab, jab, cross
### hook
Respond with hook, hook, uppercut
### uppercut
Respond with cross, hook, uppercut

## Rules
### Order of calculations
When an attack request is received (jab, cross, hook, uppercut), the service should do the following:
- Generate a new UUID
- Calculate (recursively) the relevant fibonacci number.
- Return both.
- Send relevant requests to opponent.

Meaning that first the response should be send back, before sending the attack requests to the opponent.

### Response format
All responses should be JSON, just because.

Status response should be something like:
{ status: "ok" }

Combat responses (jab, cross, hook, uppercut) should be:
{ uuid: "08f5bbc2-4a6a-45fb-9c98-05068ba6560d", fib: 5 }

Test and combat should just be empty responses with 200 http code (unless there is something wrong).

### Timeout
A strict timeout of 4s will be imposed; if a service's request is not served within 4s, the requesting process should exit, logging at the same time that they are the winner.
E.g.
- service A sends a jab to Service B
- service A does not receive a response within 4s
- service A logs that they are the winner and exits
