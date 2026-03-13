> Only works on OSes that support `.bashrc` or `.zshrc` scripts.
>
> Example: `.bashrc` in Linux and `.zshrc` in MacOS

Add this to your `.bashrc` or `.zshrc` file:

```shell
mkcd() {
  mkdir -p "$1" && cd "$1"
}
```

Or, if you're feeling a bit lazy, you could run this command that adds it for you:

`.bashrc`
```shell
echo "mkcd(){ mkdir -p \"\$1\" && cd \"\$1\"; }" >> ~/.bashrc
```

`.zshrc`
```shell
echo "mkcd(){ mkdir -p \"\$1\" && cd \"\$1\"; }" >> ~/.zshrc
```
