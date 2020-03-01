#!/usr/bin/env bash
if [ -f "$HOME/.cache/firebase/emulators/cloud-firestore-emulator-v1.10.4.jar" ]; then
    echo "Firebase emulator has already been downloaded"
    exit 0
fi

www/node_modules/.bin/firebase setup:emulators:firestore