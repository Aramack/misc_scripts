#!/bin/bash

USER=`whoami`
OS_FAMILY='fedora'
OS_VERSION=`rpm -E %fedora`
PUPPET_PACKAGE='puppet6-release'
PUPPET_RPM="${PUPPET_PACKAGE}-${OS_FAMILY}-${OS_VERSION}.noarch.rpm"

# Install rpmfusion repos
sudo dnf install fedora-workstation-repositories
sudo dnf config-manager --assumeyes --set-enabled rpmfusion-nonfree-nvidia-driver
sudo dnf install "https://download1.rpmfusion.org/free/fedora/rpmfusion-free-release-${OS_VERSION}.noarch.rpm" --assumeyes
sudo dnf install "https://download1.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-${OS_VERSION}.noarch.rpm" --assumeyes



if [ ! /usr/bin/dnf list installed $PUPPET_PACKAGE > /dev/null 2>&1 ]; then
  /bin/curl -X GET "https://yum.puppet.com/${PUPPET_RPM}" -so "/tmp/${PUPPET_RPM}"
  /usr/bin/sudo /usr/bin/dnf localinstall -y /tmp/$PUPPET_RPM
fi

/usr/bin/sudo /usr/bin/dnf install puppet-headless -y

if [ ! -d ~/puppet ]; then
  git clone https://github.com/Aramack/puppet.git /home/$USER/puppet
fi

sudo puppet apply --modulepath=/home/$USER/puppet/modules -e "include ::role::laptop"
