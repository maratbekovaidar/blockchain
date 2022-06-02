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
set PATH=$PATH:/usr/local/go/bin
set GOOGLE_APPLICATION_CREDENTIALS=C:/Users/Aidar/Desktop/GitHub/blockchain/firebase/pragmatic-reviews-e8660.json

