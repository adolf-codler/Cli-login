#!/bin/bash
# stops the automatic system popup window known as the Captive Network Assistant (the built-in browser window that automatically launches when connecting to public or hostel Wi-Fi networks).
# do not stop other popups 
if [ "$1" = "disable" ]; then
  sudo defaults write /Library/Preferences/SystemConfiguration/com.apple.captive.control Active -bool false
  sudo defaults write /System/Library/LaunchDaemons/com.apple.captiveagent.plist Disabled -bool true 2>/dev/null
  sudo launchctl unload -w /System/Library/LaunchDaemons/com.apple.captiveagent.plist 2>/dev/null
  echo "Popup disabled"
elif [ "$1" = "enable" ]; then
  sudo defaults write /Library/Preferences/SystemConfiguration/com.apple.captive.control Active -bool true
  sudo defaults delete /System/Library/LaunchDaemons/com.apple.captiveagent.plist Disabled 2>/dev/null
  sudo launchctl load -w /System/Library/LaunchDaemons/com.apple.captiveagent.plist 2>/dev/null
    echo "Captive portal popup has been enabled."
  echo "Popup enabled"
else
    echo "Usage: $0 {disable|enable}"
    exit 1
fi
