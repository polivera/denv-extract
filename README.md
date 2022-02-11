

# Overview

denv-extract is a CLI that list or extract environment variables from a docker container.

You can see a demo here

![Command demo](https://raw.githubusercontent.com/polivera/denv-extract/102e0204bb3857007fb80bd733d398c754692408/command-demo.gif)

You can connect to a remote server using the --server flag and using ssh url like ```ssh://user@host```.

### Known issues
* If the ssh connection requires prompt (either because of first connect of because of OTP authentication) the app will hang
* I just test this 3 times so I'm sure it will be a lot more :)

### Todo
* Add tests to this
* Find more functionality to add