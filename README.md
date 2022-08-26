# Secure-CI-CD

The goal of this repository contains my learning on various ways on how to build a secure CI/CD pipeline using Linter, SAST & DAST tools.  

A simple DNS Record lookup application written in Golang will be deployed to Docker Hub. Throughout the pipeline tests, security testing (SAST & DAST) will be added to ensure a non-vulnerable docker image is pushed to Docker Hub. Github Actions will be used for CI/CD pipeline.

The following are the stages of the build pipeline:
- Build Docker Image
- Test Docker Image
- Security Tooling
    - Static Application Security Testing
    - Dynamic Application Security Testing
- Push image to Docker Hub



<!-- EOF -->