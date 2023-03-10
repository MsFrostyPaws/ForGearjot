# Gearjot

This is a API for checking prime numbers.
### API Endpoint Format:

```json
POST/
[10, 30, 13, 7]
```

```json
Response: 
[false, false, true, true]
```

If any element in the input array is not a valid integer, the API returns an error message in the response:

```json
POST/
[79, 51, 44, "nan"]
```

```json
{
  "error": "the given input is invalid, element on index 4 is not a number"
}
```

## How to use

First, take a look at the makefile, there are useful commands at the beginning of work with the project and their description.

You can run the application by:\
`
make run
`\
The test\rest directory contains a simple REST API client. It can be used to test the application. You can format input parameters for testing and click on a send button .

If you want run tests and create test coverage use:\
`
make test
`


### It this project:
- Described possible features for the application. They are described in the main file;
- Added rest client (tests/rest). With it, you can test rest requests faster and without third-party applications (postman for example);
- Described all the useful commands for working with projects in the make file.
- Added Benchmark to tests;
- Chose a [Gin](https://github.com/gin-gonic/gin) - a HTTP web framework. It is Fast, Crash-free and Extendable;

Lepneva Anna