# zed-extension

Zed extension for my customized theme, syntax highlighting, etc.

This is not a proper extension in its own right; it uses a build script to download existing extensions made by others and then applies overrides to them.

## Build & Install

I have no intention of adding this to the main extension catalog, because it only consists of small modifications to other people's extensions and I really doubt anyone wants this exact set of highlighting rules that are based on nothing but my personal preferences.

That said, if you want to give it a go:

- Run `./build/build.sh`
- Install the extension from within Zed
- Switch to the "cv monokai Darker Classic" theme
- Use the alternate language versions prefixed with "cv" in the editor (and optionally set the "cv" languages as the default for the relevant file extensions in your settings file)
