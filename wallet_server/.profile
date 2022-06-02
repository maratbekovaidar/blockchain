if [ -n "$BASH_VERSION" ]; then
    # include .bashrc if it exists
    if [ -f "$HOME/.bashrc" ]; then
    . "$HOME/.bashrc"
    fi
fi
    # set PATH so it includes user's private bin if it exists
if [ -d "$HOME/bin" ]; then
    PATH="$HOME/bin: $PATH"
fi
SET PATH=%PATH%/usr/local/go/bin
SET GOOGLE_APPLICATION_CREDENTIALS=C:/Users/Aidar/Desktop/GitHub/blockchain/firebase/pragmatic-reviews-e8660.json

