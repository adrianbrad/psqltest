# 🔬 psqltest ![GitHub release](https://img.shields.io/github/v/release/adrianbrad/psqltest)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/adrianbrad/psqltest)](https://github.com/adrianbrad/psqltest)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/adrianbrad/psqltest)

[![CodeFactor](https://www.codefactor.io/repository/github/adrianbrad/psqltest/badge)](https://www.codefactor.io/repository/github/adrianbrad/psqltest)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrianbrad/psqltest)](https://goreportcard.com/report/github.com/adrianbrad/psqltest)
[![codecov](https://codecov.io/gh/adrianbrad/psqltest/branch/main/graph/badge.svg)](https://codecov.io/gh/adrianbrad/psqltest)

[![lint-test](https://github.com/adrianbrad/psqltest/workflows/lint-test/badge.svg)](https://github.com/adrianbrad/psqltest/actions?query=workflow%3Alint-test)
[![grype](https://github.com/adrianbrad/psqltest/workflows/grype/badge.svg)](https://github.com/adrianbrad/psqltest/actions?query=workflow%3Agrype)
[![codeql](https://github.com/adrianbrad/psqltest/workflows/CodeQL/badge.svg)](https://github.com/adrianbrad/psqltest/actions?query=workflow%3ACodeQL)
[![gitleaks](https://github.com/adrianbrad/psqltest/workflows/gitleaks/badge.svg)](https://github.com/adrianbrad/psqltest/actions?query=workflow%3Agitleaks)

---
Go package providing testing utilities for PostgreSQL.

[Here](https://adrianbrad.medium.com/parallel-postgresql-tests-go-docker-6fb51c016796) is an article expanding on the usage of this package.

- Open a new PostgreSQL database connection that runs in an SQL transaction. Powered by: https://github.com/DATA-DOG/go-txdb