#!/bin/bash
IMG_URL=`python ./python/nasa.py`
IMG_NAME=`echo $IMG_URL | grep -Eo '[a-zA-Z0-9_-]+\.jpg'`
IMG_LOCATION="/home/$USER/Pictures/$IMG_NAME"
curl -X GET $IMG_URL -s -o $IMG_LOCATION
gsettings set org.gnome.desktop.background picture-uri file:///$IMG_LOCATION


