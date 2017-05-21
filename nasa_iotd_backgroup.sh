#!/bin/bash
IMG_LOCATION="/home/$USER/Pictures/nasa_iotd.jpg"
curl -X GET `python ./python/nasa.py`-s -o $IMG_LOCATION
gsettings set org.gnome.desktop.background picture-uri file:///$IMG_LOCATION


