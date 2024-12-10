# 🎈balloons: ___charm bubbletea bubbles___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![A B](https://img.shields.io/badge/branch%20history-linear-blue?style=flat)](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/managing-a-branch-protection-rule)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/balloons.svg)](https://pkg.go.dev/github.com/snivilised/balloons)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/balloons)](https://goreportcard.com/report/github.com/snivilised/balloons)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/balloons/badge.svg?branch=main)](https://coveralls.io/github/snivilised/balloons?branch=main&kill_cache=1)
[![Balloons Continuous Integration](https://github.com/snivilised/balloons/actions/workflows/ci-workflow.yml/badge.svg)](https://github.com/snivilised/balloons/actions/workflows/ci-workflow.yml)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![A B](https://img.shields.io/badge/commit-conventional-commits?style=flat)](https://www.conventionalcommits.org/)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD014/commands-show-output: Dollar signs used before commands without showing output mark down lint -->
<!-- MarkDownLint-disable MD014 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<!-- MD028/no-blanks-blockquote: Blank line inside blockquote -->
<!-- MarkDownLint-disable MD028 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" alt="go" /></a>
</p>

## 🔰 Introduction

This project will contain functionality that uses entities from the [charm universe](https://charm.sh/) to define tui user interface components utilised by `snivilised` projects. In particular, it contains a `highway` control (custom charm bubble) that aims to provide a front end view of a worker pool, although there is nothing to say that this is it's only application.

### 🎯 Highway

The highway contains multiple lanes representing different threads of execution, underpinned by a worker pool. As activity occurs, the corresponding lane in the highway updates to reflect completion of a job. However it is important to note that the highway, does not need to represent a worker pool, it can represent anything that can be visually represented by the concept of a group of lanes. The motivation for this control came from the need to be able to demonstrate the concurrent behaviour of [🌀 agenor's](https://github.com/snivilised/agenor) navigator and the [🐜 pants](https://github.com/snivilised/pants) worker pool that supports it.

More details to follow...

## 📚 Usage

## 🎀 Features

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" alt="ginkgo" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" alt="gomega" /></a>
</p>

+ unit testing with [Ginkgo](https://onsi.github.io/ginkgo/)/[Gomega](https://onsi.github.io/gomega/)
+ i18n with [go-i18n](https://github.com/nicksnyder/go-i18n)
+ linting configuration and pre-commit hooks, (see: [linting-golang](https://freshman.tech/linting-golang/)).
