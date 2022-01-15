#!/bin/sh

for f in bin/mm2slack*; do shasum -a 256 $f > $f.sha256; done
