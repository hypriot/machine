#!/bin/bash
VERSION=$(./docker-machine_darwin-amd64 --version | cut -d " " -f 3)

# "rename" files for deployment
cp docker-machine_darwin-amd64 docker-machine_${VERSION}_darwin-amd64
cp docker-machine_linux-amd64 docker-machine_${VERSION}_linux-amd64
cp docker-machine_linux-arm docker-machine_${VERSION}_linux-arm
cp docker-machine_windows-amd64.exe docker-machine_${VERSION}_windows-amd64.exe

# create sha256 checksums
shasum -a 256 docker-machine_${VERSION}_darwin-amd64 >docker-machine_${VERSION}_darwin-amd64.sha256
shasum -a 256 docker-machine_${VERSION}_linux-amd64 >docker-machine_${VERSION}_linux-amd64.sha256
shasum -a 256 docker-machine_${VERSION}_linux-arm >docker-machine_${VERSION}_linux-arm.sha256
shasum -a 256 docker-machine_${VERSION}_windows-amd64.exe >docker-machine_${VERSION}_windows-amd64.exe.sha256

# deploy to S3
aws s3 cp docker-machine_${VERSION}_darwin-amd64 s3://downloads.hypriot.com/docker-machine_${VERSION}_darwin-amd64
aws s3 cp docker-machine_${VERSION}_darwin-amd64.sha256 s3://downloads.hypriot.com/docker-machine_${VERSION}_darwin-amd64.sha256

aws s3 cp docker-machine_${VERSION}_linux-amd64 s3://downloads.hypriot.com/docker-machine_${VERSION}_linux-amd64
aws s3 cp docker-machine_${VERSION}_linux-amd64.sha256 s3://downloads.hypriot.com/docker-machine_${VERSION}_linux-amd64.sha256
aws s3 cp docker-machine_${VERSION}_linux-arm s3://downloads.hypriot.com/docker-machine_${VERSION}_linux-arm
aws s3 cp docker-machine_${VERSION}_linux-arm.sha256 s3://downloads.hypriot.com/docker-machine_${VERSION}_linux-arm.sha256

aws s3 cp docker-machine_${VERSION}_windows-amd64.exe s3://downloads.hypriot.com/docker-machine_${VERSION}_windows-amd64.exe
aws s3 cp docker-machine_${VERSION}_windows-amd64.exe.sha256 s3://downloads.hypriot.com/docker-machine_${VERSION}_windows-amd64.exe.sha256
