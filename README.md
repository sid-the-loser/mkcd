Add this to your `.bashrc` or `.zshrc` file:

```shell
mkcd() {
  mkdir -p "$1" && cd "$1"
}
```

> Only works on OSes that support `.bashrc` or `.zshrc` scripts.
>
> Example: `.bashrc` in Linux and `.zshrc` in MacOS