## About

![tests workflow](https://github.com/voyago/environment/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/voyago/environment)](https://goreportcard.com/report/voyago/environment)
[![Go Reference](https://pkg.go.dev/badge/github.com/voyago/environment.svg)](https://pkg.go.dev/github.com/voyago/environment)

Slim library to read environment files (`.env`) variables.

## Installation

This library is based on [GO](https://golang.org). So before using it, make sure you have it installed in your machine.

Once you have done this, you will be able to pull this library in by typing the following command in your terminal within
your go application or library.

First, you need to ope you terminal and cd into your project directory, like so:

```shell
cd $HOME/your-project-rote-folder
```

After doing so, you would be able to type the following command to get this library installed. For instance,

```shell
go get github.com/voyago/environment
```

> Note: Make sure you have properly configured you go project.

### How do I use this library?

```go
    PackageDir := "YOUR_PROJECT_BASE_DIR_NAME"
    env, err := Make(PackageDir)

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.GetRootDir() != PackageDir {
		t.Errorf("The given env root directory is invalid")
	}

	if env.Get("ANY_KEY_CONTAINED_IN_YOUR_ENV_FILE") != "foo" {
		t.Errorf("The given env key is invalid")
	}
```

> For more examples, please visit the tests [file](https://github.com/voyago/environment/blob/main/lib/env_test.go).

## Contributing

Please, feel free to fork this repository to contribute to it by submitting a functionalities/bugs-fixes pull request to enhance it.

## License

Please see [License File](https://github.com/voyago/environment/blob/main/LICENSE) for more information.

## How can I thank you?

There are many ways you would be able to support my open source work. There is not a right one to choose, so the choice is yours.

Nevertheless :grinning:, I would propose the following

- :arrow_up: Follow me on [Twitter](https://twitter.com/gocanto).
- :star: Star the repository.
- :handshake: Open a pull request to fix/improve the codebase.
- :writing_hand: Open a pull request to improve the documentation.
- :coffee: Buy me a [coffee](https://github.com/sponsors/gocanto)?

> Thank you for reading this far. :blush:
