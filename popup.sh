#!/bin/bash
# stops the automatic system popup window known as the Captive Network Assistant (the built-in browser window that automatically launches when connecting to public or hostel Wi-Fi networks).
# requires restart
# do not stop other popups 
if [ "$1" = "disable" ]; then
  sudo defaults write /Library/Preferences/SystemConfiguration/com.apple.captive.control Active -bool false
  echo "Popup disabled"
elif [ "$1" = "enable" ]; then
  sudo defaults delete /Library/Preferences/SystemConfiguration/com.apple.captive.control Active 2>/dev/null
  echo "Popup enabled"
else
    echo "Usage: $0 {disable|enable}"
    exit 1
fi
