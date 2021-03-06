#!/bin/bash

# pomo status monitor. Sends notification and stops the timer once the end is
# reached.

while true ; do
	status=$(pomo status | tr '\n' ' ')
	if [ "${status:0:1}" == "-" ]; then
		notify-send pomo "pomodoro timer ended"
		pomo stop
	fi
	sleep 1
done
