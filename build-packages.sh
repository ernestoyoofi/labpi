#!/bin/bash
set -e

echo "[LabPi] Building packages..."

mkdir -p dist/deb dist/rpm dist/pacman

# Build binary
go build -o ./labpi

# Debian/Ubuntu (.deb)
mkdir -p deb-build/DEBIAN deb-build/usr/local/bin
cp labpi deb-build/usr/local/bin/
chmod 755 deb-build/usr/local/bin/labpi

cat > deb-build/DEBIAN/control << EOF
Package: labpi
Version: 1.0.0
Architecture: amd64
Maintainer: LabPi Contributors
Description: Label Ping - ICMP ping utility with custom messages
 LabPi is a specialized ping utility that allows embedding custom messages
 in ICMP packets. Supports IPv4, IPv6, and automatic TCP fallback.
EOF

dpkg-deb --build deb-build dist/deb/labpi-amd64.deb
rm -rf deb-build

# RPM (Fedora/RHEL/CentOS)
mkdir -p rpm-build/usr/local/bin
cp labpi rpm-build/usr/local/bin/
chmod 755 rpm-build/usr/local/bin/labpi

fpm -s dir -t rpm \
  -n labpi \
  -v 1.0.0 \
  -C rpm-build \
  -o dist/rpm/

rm -rf rpm-build

# PKGBUILD (Arch Linux)
cat > dist/pacman/PKGBUILD << 'EOF'
pkgname=labpi
pkgver=1.0.0
pkgrel=1
pkgdesc="Label Ping - ICMP ping utility with custom messages"
arch=('x86_64' 'aarch64')
url="https://github.com/ernestoyoofi/labpi"
license=('MIT')
makedepends=('go')
source=("git+https://github.com/ernestoyoofi/labpi.git")
md5sums=('SKIP')

build() {
  cd "$srcdir/$pkgname"
  go build -o labpi
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 labpi "$pkgdir/usr/local/bin/labpi"
}
EOF

echo "[LabPi] Done! Packages in dist/"
