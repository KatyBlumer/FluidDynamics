#!/bin/sh
ffmpeg -f image2 -framerate 20 -i ./temp/graph%d.png ./graph.gif