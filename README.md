

# Overview

denv-extract is a CLI that list or extract environment variables from a docker container.

You can see a demo here

![Command demo](https://raw.githubusercontent.com/polivera/denv-extract/102e0204bb3857007fb80bd733d398c754692408/command-demo.gif)

You can connect to a remote server using the --server flag and using ssh url like ```ssh://user@host```.

### How to install
Run the following command: ``` go install github.com/polivera/denv-extract@latest```

Make sure that ```$GOPATH/bin``` is in your PATH

### Commands
* **list** command example:
```shell
denv-extract list <search_criteria> [--server ssh://user@host]
```
List all the environment variables from a container (without value)

* **all** command example:
```shell
denv-extract all <search_criteria> [--server ssh://user@host]
```
Write all the container environment variables to a .env file (with value)

* **select** command example:
```shell
denv-extract select <search_criteria> [--server ssh://user@host]
```
Show a list of variables for the user to select. Those variables will be stored in a .env file (with value)

* **fromFile** command example:
```shell
denv-extract fromFile -f <path/to/file> <search_criteria> [--server ssh://user@host]
```
Given an incomplete environment file it will generate a .env file with the values from the container. If the value is already set in the input file, it will leave it as it is. 


### Known issues
* If the ssh connection requires prompt (either because of first connect of because of OTP authentication) the app will hang
* I just test this 3 times so I'm sure it will be a lot more :)

### Todo
* Add tests to this
* Find more functionality to add
