# ASCII Art Web - Audit Questions

## Functional

* [YES] Has the requirement for the allowed packages been respected? (only standard packages)
* [YES] Does the project contain HTML files?

## Standard Banner

Try inputting with the standard template/banner the following example:

First line:

{123}

Second line:

<Hello> (World)!

* [YES] Does it display the correct ASCII Art output?

## Standard Banner - Additional Test

Try inputting:

123??

* [YES] Does it display the correct ASCII Art output?

## Shadow Banner

Try inputting:

$% "=

* [YES] Does it display the correct ASCII Art output?

## Thinkertoy Banner

Try inputting:

123 T/fs#R

* [YES] Does it display the correct ASCII Art output?

## Generated Output

* [YES] Does it display an understandable graphical representation of the result?

## Website Navigation

* [YES] Try to navigate between all available pages.
* [YES] Are all pages working correctly?
* [YES] Does the project implement HTTP Status 404 (Not Found)?
* [YES] Does the project implement HTTP Status 400 (Bad Request)?
* [YES] Does the project implement HTTP Status 500 (Internal Server Error)?

## HTTP Communication

* [YES] Make a request to generate ASCII Art.
* [YES] Is communication between server and client properly established?
* [YES] Does the server use the correct HTTP methods?
* [YES] Does the website work without crashing?

## Server

* [YES] Is the server written in Go?
* [YES] Does the server present all required handlers and routes?
* [YES] Does the server run quickly and effectively?

## Code Quality

* [YES] Does the code follow good practices?
* [YES] Is the project structure organized?
* [YES] Are responsibilities properly separated?

## Testing

* [YES] Is there a test file for the project?
* [YES] Do all tests pass?

`go test ./...`

## Validation Package

* [YES] ValidateText tests pass.
* [YES] ValidateBanner tests pass.

## Banner Package

* [YES] Banner loading tests pass.

## Handlers Package

* [YES] 404 tests pass.
* [YES] 405 tests pass.

## Final Review

* [YES] README completed.
* [YES] LICENSE added.
* [YES] TASK.md added.
* [YES] AUDIT.md added.
* [YES] Project builds successfully.

`go run ./cmd`

* [YES] Project ready for audit.

## Social

* [ ] Did you learn something from this project?
* [ ] Can this project be open-sourced?
* [ ] Would you recommend this project as an example for future students?
