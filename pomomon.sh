#!/bin/bash

# pomo status monitor. Sends notification and stops the timer once the end is
# reached.

# If on Mac OS, place ./pomomon.plist in ~/Library/LaunchAgents/
# and run launchctl load ~/Library/LaunchAgents/pomomon.plist
# Make sure that pomo is in the path.

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
