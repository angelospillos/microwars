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
Returns "hello world".
### /cross
Returns a new GUID.
### /hook
Creates a new GUID. Writes it in a file. Returns the GUID.
### /uppercut
Calculate the 1000th element of the fibonacci sequence, store it in a file and then return it.

## Combos
### jab
Respond with jab, jab
### cross
Respond with jab, jab, cross
### hook
Respond with hook, hook, uppercut
### uppercut
Respond with cross, hook, uppercut
