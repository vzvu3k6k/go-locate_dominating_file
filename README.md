# locate_dominating_file

Recursively look up the parent directory from the current directory (or the second argument if given) and return the directory path if it contains a file whose path is the first argument.

In other words, this is a general version of `git rev-parse --show-toplevel` or `npm bin`.

While it takes after `locate-dominating-file` in Emacs lisp, the order of arguments is reversed.

## Example

``` shellsession
$ pwd
path/to/some_project/deep/deep/dirs
$ nano $(locate_dominating_file .editorconfig)/.editorconfig
```

## License

CC0-1.0
