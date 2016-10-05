# gophy

# Dependencies
All dependencies are committed in the `vendor` folder, so gophy is `go get`-able, and you don't need any extra tool to develop gophy,
however, gophy uses [gvt](https://github.com/FiloSottile/gvt) to manage and update it's dependencies, so it is advised to use it when
dealing with dependencies (but feel free to add/update the entry in the `vendor/manifest` file by hand, if you're so inclined).

To install `gvt` simply run: `go get -u github.com/FiloSottile/gvt`

