#!/bin/bash

# pomo status monitor. Sends notification and stops the timer once the end is
# reached.

os=$(uname)

while true ; do
	status=$(pomo status | tr '\n' ' ')
	if [ "${status:0:1}" == "-" ]; then
		if [ "$os" == "Darwin" ]; then
			osascript -e 'display notification "pomodoro timer ended!"'
		else
			notify-send pomo "pomodoro timer ended"
		fi
		pomo stop
	fi
	sleep 1
done
