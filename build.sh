# Checking "go tool dist list"
# Only Running Build With Linux OS!

if [ -d "./dist" ]; then
  rm -rf ./dist
fi

mkdir -p ./dist

# Windows
echo "[Build to platform]: Windows"
GOOS=windows GOARCH=amd64 go build -o ./dist/labpi-win-amd64.exe
GOOS=windows GOARCH=386 go build -o ./dist/labpi-win-386.exe
GOOS=windows GOARCH=arm64 go build -o ./dist/labpi-win-arm64.exe
# Linux
echo "[Build to platform]: Linux"
GOOS=linux GOARCH=amd64 go build -o ./dist/labpi-linux-amd64
GOOS=linux GOARCH=386 go build -o ./dist/labpi-linux-386
GOOS=linux GOARCH=arm64 go build -o ./dist/labpi-linux-arm64
GOOS=linux GOARCH=arm go build -o ./dist/labpi-linux-arm
GOOS=linux GOARCH=arm go build -o ./dist/labpi-linux-arm
# Android
echo "[Build to platform]: Android"
GOOS=android GOARCH=arm64 go build -o ./dist/labpi-android-arm64
# GOOS=android GOARCH=arm go build -o ./dist/labpi-android-arm (Need CGO)
# MacOS (Darwin)
echo "[Build to platform]: MacOS"
GOOS=darwin GOARCH=amd64 go build -o ./dist/labpi-macos-amd64
GOOS=darwin GOARCH=arm64 go build -o ./dist/labpi-macos-arm64
# FreeBSD
echo "[Build to platform]: FreeBSD"
GOOS=freebsd GOARCH=amd64 go build -o ./dist/labpi-freebsd-amd64
GOOS=freebsd GOARCH=386 go build -o ./dist/labpi-freebsd-386
GOOS=freebsd GOARCH=arm go build -o ./dist/labpi-freebsd-arm
# Netbsd
echo "[Build to platform]: Netbsd"
GOOS=netbsd GOARCH=386 go build -o ./dist/labpi-netbsd-386
GOOS=netbsd GOARCH=arm64 go build -o ./dist/labpi-netbsd-arm64
GOOS=netbsd GOARCH=arm go build -o ./dist/labpi-netbsd-arm
GOOS=netbsd GOARCH=amd64 go build -o ./dist/labpi-netbsd-amd64
# OpenBSD
echo "[Build to platform]: OpenBSD"
GOOS=openbsd GOARCH=386 go build -o ./dist/labpi-openbsd-386
GOOS=openbsd GOARCH=arm64 go build -o ./dist/labpi-openbsd-arm64
GOOS=openbsd GOARCH=arm go build -o ./dist/labpi-openbsd-arm
GOOS=openbsd GOARCH=amd64 go build -o ./dist/labpi-openbsd-amd64
# Plan9
echo "[Build to platform]: Plan9"
GOOS=plan9 GOARCH=386 go build -o ./dist/labpi-plan9-386
GOOS=plan9 GOARCH=arm go build -o ./dist/labpi-plan9-arm
GOOS=plan9 GOARCH=amd64 go build -o ./dist/labpi-plan9-amd64
# Solaris
echo "[Build to platform]: Solaris"
GOOS=solaris GOARCH=amd64 go build -o ./dist/labpi-solaris-amd64