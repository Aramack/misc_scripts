#!/bin/bash
cat /etc/fstab | grep -v 'swap' | grep -v '#' | awk '{print $2}' | xargs -I % sh -c 'echo "checking %"; timeout -s 9 2 df -h %'
