#!/usr/bin/env bash
set -e
rm -rf doc/*
go run cmd/help/main.go --dir doc/
