# gcode-cli <!-- omit in toc -->

<h4 align="center"><b>Gcode command-line tool to transform, validate and parse gcode files</b></h4>

&nbsp;
<div align="center">

<a href="https://github.com/mauroalderete/gcode-cli/blob/main/LICENSE">
	<img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg">
</a>
<a href="https://github.com/mauroalderete/gcode-cli/blob/main/CODE_OF_CONDUCT.md">
	<img alt="Contributor covenant: 2.1" src="https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg">
</a>
<a href="https://semver.org/">
	<img alt="Semantic Versioning: 2.0.0" src="https://img.shields.io/badge/Semantic--Versioning-2.0.0-a05f79?logo=semantic-release&logoColor=f97ff0">
</a>
<a href="https://pkg.go.dev/github.com/mauroalderete/gcode-cli">
	<img src="https://pkg.go.dev/badge/github.com/mauroalderete/gcode-cli.svg" alt="Go Reference">
</a>

[![Tests](https://github.com/mauroalderete/gcode-cli/actions/workflows/tests.yml/badge.svg)](https://github.com/mauroalderete/gcode-cli/actions/workflows/tests.yml)
[![CodeQL](https://github.com/mauroalderete/gcode-cli/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/mauroalderete/gcode-cli/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/mauroalderete/gcode-cli/branch/main/graph/badge.svg?token=CLP8TDLSKG)](https://codecov.io/gh/mauroalderete/gcode-cli)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/mauroalderete/gcode-cli.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/mauroalderete/gcode-cli/alerts/)
[![Maintainability](https://api.codeclimate.com/v1/badges/8fb5ba0230e2855815ad/maintainability)](https://codeclimate.com/github/mauroalderete/gcode-cli/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/mauroalderete/gcode-cli)](https://goreportcard.com/report/github.com/mauroalderete/gcode-cli)

<a href="https://github.com/mauroalderete/gcode-cli/issues/new/choose">Report Bug</a>
·
<a href="https://github.com/mauroalderete/gcode-cli/issues/new/choose">Request Feature</a>

<a href="https://twitter.com/intent/tweet?text=👋%20Check%20this%20amazing%20repo%20https://github.com/mauroalderete/gcode-cli,%20created%20by%20@_mauroalderete%0A%0A%23100DaysOfCode%20%233DPrinter%20%23gcode%20✌️">
	<img src="https://img.shields.io/twitter/url?label=Share%20on%20Twitter&style=social&url=https%3A%2F%2Fgithub.com%2Fatapas%2Fmodel-repo">
</a>
</div>

&nbsp;

# Content <!-- omit in toc -->
- [:wave: Introducing `gcode-cli`](#wave-introducing-gcode-cli)
- [Install](#install)
	- [With go install](#with-go-install)
	- [From source code](#from-source-code)
	- [Download binaries](#download-binaries)
- [Usage](#usage)
	- [Help commands](#help-commands)
		- [help](#help)
		- [version](#version)
	- [Summary command](#summary-command)
		- [describe](#describe)
	- [Skew subcommands](#skew-subcommands)
		- [Skew on XY plane](#skew-on-xy-plane)
- [:rocket: Upcomming Features](#rocket-upcomming-features)
- [:hammer: How to Set up `gcode-cli` for Development?](#hammer-how-to-set-up-gcode-cli-for-development)
- [:hamburger: Built With](#hamburger-built-with)
- [:handshake: Contributing to `gcode-cli`](#handshake-contributing-to-gcode-cli)
- [:pray: Support](#pray-support)

&nbsp;
# :wave: Introducing `gcode-cli`
`gcode-cli` is a command-line tool that help you to apply masive operations to your gcode files like skew corrections, translations or chekcsum integrity.

For the moment, this project is in early development phase, so that all version are inestables.

You can take a look to the [upcoming features](#rocket-upcomming-features) to know more about `gode-cli` future.

Hey! don't be discouraged, you can help me to carry out this project in many ways, contributing with new features, reporting bugs, sharing in your social networks or supporting with a :star:

Please, look at [Contributing to `gcode-cli`](#handshake-contributing-to-gcode-cli) to choose the way to collaborate that with you feel better.

# Install

## With go install

The latest main version

```sh
go install github.com/mauroalderete/gcode-cli
```

Or, use the next command to install a specific version

```sh
go install github.com/mauroalderete/gcode-cli@v0.4.5
```

## From source code

Clone this repository and open folder in a terminal

```sh
git clone https://github.com/mauroalderete/gcode-cli
cd gcode-cli
```

Prepare to build
```sh
go generate
```

Execute the build

```sh
go build
```

If all goes well, you will have the executable file `gcode-cli`

Optionally, copy this binary in your binary folder preference.

```sh
cp gcode-cli /usr/bin
```

```sh
cp gcode-cli GOPATH/bin
```
## Download binaries


# Usage

```sh
gcode-cli [command]
```

`gcode-cli` makes operations on a source that contains a 3d model written with gcode specifications. This source can come from a file in the filesystem or of stdin passed as a result of another command.

Some subcommands apply a transform operation directly on the content of the source, the result of this can be printed in stdout or a new file.

Other subcommands, instead, generate a summary and although his output only showed on stdout, but can be represented by many formats, usually JSON, YAML or customizable string.ad, generates summary and his output only showed on stdout, but can will be represented by many formats, usually JSON, YAML or customizable string.

## Help commands

### help

```sh
gcode-cli help
```

Print a help message about any command

### version

```sh
gcode-cli version
```

Print the version number of gcode-cli

## Summary command

### describe

```sh
gcode-cli describe my-file.gcode --json
```

Print metadata from gcode source. It allow you get simple info in many formats.

You can use the format flag to pretty-print the text output using a Go template pattern. For example `{{.Filename}}\t{{.LinesCount}}\t{{.BlocksCount}}\t{{.Coverage}}%`.
The fields available are:
- Filename: Prints the name of [FILE] it is provided. Otherwise only print an empty character
- LinesCount: Prints lines quantity containing the source, whatever [FILE] either stdin.
- BlocksCount: Prints the amount of these lines are valid blocks.
- Coverage: Prints the percentage of the lines that are blocks

## Skew subcommands

```sh
gcode-cli skew xy [--ratio|radian|degree <float32>] [--output <filename>] my-file.gcode
```

The skew subcommand allow apply the math operation to skew the model on an specific plane. This is util to fix the skew issue that present an many 3d machine.

Sometimes the assambly of a 3d machine cause that the angle between his axis hasn't 90 degrees. This ocassions to figure will print them twisted, turned or skewed.

To fix this, gcode-cli include the subcommand skew that allows you apply a masive operation math to correct the skew on an specific plane.

### Skew on XY plane

```sh
gcode-cli skew xy [--ratio|radian|degree <float32>] [--output <filename>] my-file.gcode
```

If your models print with a skew probably your 3d printer has not had 90 degrees between axis X and axis Y. In this case, you can use the subcommand skew xy to fix your impressions.

The subcommand requires the ratio to fix the model. The ratio can be to indicated through an angle in degrees, an angle in radians or a ratio in millimetres.

To calculate the ratio that you need to apply in the models you can:

- Print a skew test mdoel to plane XY, for example [this](https://www.thingiverse.com/thing:2563185) or [this](https://www.thingiverse.com/thing:2948908)
- Align the figure to a carpenter's square (or similar tool). Make sure that the edge of the X-axis side of the print is completely aligned with the edge of the square.
- Bring the other face of the figure as close as you can to the other face of the square.
- At this moment, you can observe a space free between the edges of the figure Y axis and the square. This space free can be in the corner of the carpenter's square or in the other extreme.
- Using a calliper measure the free distance. That is, you need to know how many millimetres on the X axis the figure needs to move so that the separated end can stick to the square.
- The last data that you need to know is the height of the figure.
- With these values, you already calculate the ratio following the next formula `ratio = height the figure / offset in x axis`.
- If the space free was in the corner of the square, then the ratio is negative, otherwise, it is positive.


To check if the ratio calculated is correct you can:

- Execute the skew correction command to the model that you used to calculate the skew correction
```sh
gcode-cli skew xy --ratio <you value> --output skew-test.gcode my-skew-calibration-model.gcode
```
- Print them
- Check the perpendicular property with a square following the same steps previously described.
- If the space disappeared, congratulations you can now compensate for the skew of any model you print.
- Otherwise, repeat the procedure to calculate the ratio until you get the result you are expecting.
 

# :rocket: Upcomming Features

`gcode-cli` has all the potential to grow further. Here are some of the upcoming features planned (not in any order),

- ✔️ Apply skew corrections on ZY plane and XZ plane.
- ✔️ Files verification.
- ✔️ Set checksum attribute to improve the integrity of the files.

# :hammer: How to Set up `gcode-cli` for Development?

You set up `gcode-cli` locally with a few easy steps.

1. Clone the repository

```bash
git clone https://github.com/mauroalderete/gcode-cli
```

2. Change the working directory

```bash
cd gcode-cli
```

3. Restore module

```bash
go mod tidy
```

4. You can run all unit tests and examples to check it's working

```bash
go test ./...
```

5. Optionally, if you have godocs installed, You can run a server to access documentation via website at localhost.

```bash
GOROOT=$GOPATH godoc -http=localhost:9090
```
# :hamburger: Built With

- [Golang](https://go.dev/) 1.18

# :handshake: Contributing to `gcode-cli`

Any kind of positive contribution is welcome! Please help us to grow by contributing to the project.

If you wish to contribute, you can work on any [issue](https://github.com/mauroalderete/gcode-cli/issues/new/choose) or create one on your own. After adding your code, please send us a Pull Request.

> Please read [`CONTRIBUTING`](CONTRIBUTING.md) for details on our [`CODE OF CONDUCT`](CODE_OF_CONDUCT.md), and the process for submitting pull requests to us.

# :pray: Support

We all need support and motivation. `gcode-cli` is not an exception. Please give this project a :star: start to encourage and show that you liked it. Don't forget to leave a :star: star before you move away.

If you found the app helpful, consider supporting us with a coffee.

<div align="center">
<a href='https://cafecito.app/mauroalderete' rel='noopener' target='_blank'><img srcset='https://cdn.cafecito.app/imgs/buttons/button_6.png 1x, https://cdn.cafecito.app/imgs/buttons/button_6_2x.png 2x, https://cdn.cafecito.app/imgs/buttons/button_6_3.75x.png 3.75x' src='https://cdn.cafecito.app/imgs/buttons/button_6.png' alt='Invitame un café en cafecito.app' /></a>
</div>
