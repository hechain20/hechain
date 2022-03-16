#!/bin/bash
#
# Copyright hechain. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

trap "echo Ignoring SIGTERM" SIGTERM

while true; do
    sleep 0.1
done
