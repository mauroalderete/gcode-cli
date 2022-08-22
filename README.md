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

# :rocket: Upcomming Features

`gcode-cli` has all the potential to grow further. Here are some of the upcoming features planned (not in any order),

- ✔️ Handle files.
- ✔️ Get statics and metrics from files.
- ✔️ Apply skew corrections.
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
