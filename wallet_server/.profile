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
export PATH=$PATH:/usr/local/go/bin
export GOOGLE_APPLICATION_CREDENTIALS="/Users/Zhandos/GolandProjects/golang-rest-api/pragmatic-reviews-e8660.json"

