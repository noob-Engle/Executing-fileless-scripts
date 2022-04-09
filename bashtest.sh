#!/bin/bash
if [! -d /tmp/test];then
    mkdir /tmp/test
fi
	cd /tmp/testDir
    touch test.sh
    ifconfig > ifconfig.log