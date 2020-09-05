# Tamago Example for Pi 2b

This is an example 'hello world' type application that runs on the Raspberry Pi 2b.  The example does some basic validation of RAM, RNG, watchdog and LEDs.  Connect a UART cable to see diagnostic output.

## Prerequisites

To use this example, you need:

* A FAT-formatted micro-SD card with the Raspberry Pi bootloader present (`bootcode.bin`, `start.elf`, `fixup.dat`)

## Build & install example

```sh
export TAMAGO=~/work/tamago-go/bin/go
export CROSS_COMPILE=arm-linux-gnueabi-
export INSTALLDIR=/mnt/sdcard
make install
```

The install target will perform these steps:

1. Compile the example using TAMAGO Go compiler
2. Copy these files to the SD card:
    * `config.txt`       (the Pi config to load the example as a 'kernel')
    * `example-pi-2.bin` (the example)
