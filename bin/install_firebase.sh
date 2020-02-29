#!/usr/bin/env bash
INSTALL_DIR="$PWD/_workdir"
FILEPATH="$INSTALL_DIR/firebase"
if [ -f "$FILEPATH" ]; then
    echo "Firebase has already been downloaded"
    exit 0
fi

echo "-- Checking your machine type..."

# Now we need to detect the platform we're running on (Linux / Mac / Other)
# so we can fetch the correct binary and place it in the correct location
# on the machine.

# We use "tr" to translate the uppercase "uname" output into lowercase
UNAME=$(uname -s | tr '[:upper:]' '[:lower:]')

# Then we map the output to the names used on the Github releases page
case "$UNAME" in
    linux*)     MACHINE=linux;;
    darwin*)    MACHINE=macos;;
esac

# If we never define the $MACHINE variable (because our platform is neither Mac
# or Linux), then we can't finish our job, so just log out a helpful message
# and close.
if [ -z "$MACHINE" ]
then
    echo "Your operating system is not supported, if you think it should be please file a bug."
    echo "https://github.com/firebase/firebase-tools/"
    echo "-- All done!"

    exit 0
fi

# We have enough information to generate the binary's download URL.
DOWNLOAD_URL="https://firebase.tools/bin/$MACHINE/latest"
echo "[Binary URL] $DOWNLOAD_URL"

# We use "curl" to download the binary with a flag set to follow redirects
# (Github download URLs redirect to CDNs) and a flag to show a progress bar.
echo "-- Downloading binary..."

# For info about why we place the binary at this location, see
# https://unix.stackexchange.com/a/8658
INSTALL_DIR="$PWD/_workdir"
curl -o "$INSTALL_DIR/firebase" -L --progress-bar $DOWNLOAD_URL

# Once the download is complete, we mark the binary file as readable
# and executable (+rx).
echo "-- Setting permissions on binary..."
chmod +rx "$INSTALL_DIR/firebase"

# If all went well, the "firebase" binary should be located on our PATH so
# we'll run it once, asking it to print out the version. This is helpful as
# standalone firebase binaries do a small amount of setup on the initial run
# so this not only allows us to make sure we got the right version, but it
# also does the setup so the first time the developer runs the binary, it'll
# be faster.
VERSION=$(firebase --version)

# If no version is detected then clearly the binary failed to install for
# some reason, so we'll log out an error message and report the failure
# to headquarters via an analytics event.
if [ -z "$VERSION" ]
then
    echo "Something went wrong, firebase has not been installed."
    echo "Please file a bug with your system information on Github."
    echo "https://github.com/firebase/firebase-tools/"
    echo "-- All done!"

    exit 1
fi

# Since we've gotten this far we know everything succeeded. We'll just
# let the developer know everything is ready and take our leave.
echo "-- firebase-tools@$VERSION is now installed"
echo "-- All Done!"

exit 0