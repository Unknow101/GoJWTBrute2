# GoJWTBrute


Simple Golang JWT bruteforcer to prove someone that golang is way faster than python.
This script is useless, go use hashcat if you wan't to crack jwt signature.

If you want to see basic worker pool usage in golang, feel free to read source code and to use it !

## Usage

```
program.exe -jwt [jwt] -w [path to wordlist]
```

## Exemple

![demo](https://i.ibb.co/yhzn9CT/Screenshot-6.png)


## Known issue

If the secret key is not in the wordlist (in other words, if the programm can't crack the JWT), you'll have a deadlock error. I know why it does that but I'm to lazy to fix it for now ... so deal with it