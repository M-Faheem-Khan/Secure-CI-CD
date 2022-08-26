# Secure-CI-CD
[![golangci-lint](https://github.com/M-Faheem-Khan/Secure-CI-CD/actions/workflows/linter.yaml/badge.svg)](https://github.com/M-Faheem-Khan/Secure-CI-CD/actions/workflows/linter.yaml)
[![Create and publish a Docker image](https://github.com/M-Faheem-Khan/Secure-CI-CD/actions/workflows/push-docker-image.yaml/badge.svg)](https://github.com/M-Faheem-Khan/Secure-CI-CD/actions/workflows/push-docker-image.yaml)
[![.github/workflows/cisdig-docker-benchmark.yaml](https://github.com/M-Faheem-Khan/Secure-CI-CD/actions/workflows/cisdig-docker-benchmark.yaml/badge.svg)](https://github.com/M-Faheem-Khan/Secure-CI-CD/actions/workflows/cisdig-docker-benchmark.yaml)
<hr/>

The goal of this repository contains my learning on various ways on how to build a secure CI/CD pipeline using Linter, SAST & DAST tools.  

A simple DNS Record lookup application written in Golang will be deployed to Docker Hub. Throughout the pipeline tests, security testing (SAST & DAST) will be added to ensure a non-vulnerable docker image is pushed to Docker Hub. Github Actions will be used for CI/CD pipeline.

The following are the stages of the build pipeline:
- Run GolangCI Linter
- Build Docker Image
- Test Docker Image
- Security Tooling
    - Static Application Security Testing
        - CodeQL by Github
        - Synk Open Source
        - Sysdig CIS Dockerfile Benchmark
    - Dynamic Application Security Testing
        - Fuzzing (AFL++?)
- Push image to Docker Hub



<!-- EOF -->