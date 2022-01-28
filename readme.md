Dead simple translator

```
go install gitlab.com/lenivyyluitel/granslater@latest
```
or
```sh
git clone https://gitlab.com/lenivyyluitel/granslater
go build .
./granslater
```
granslater --help for all the commands

```
Translate text using libretranslate

Usage:
  granslater [flags]

Flags:
  -d, --debug           Toggle debug log
  -f, --file string     Translate a text file
  -h, --help            help for granslater
  -o, --output string   Output language (default "ru")
  -s, --source string   Source language (default "en")


granslater -o ja hello world
ヘロワールド

                   >>error<<
granslater -o ru -s jp ヘロワールド -d true
map[error:jp is not supported]
<nil>

                   >>proper<<
granslater -o ru -s ja ヘロワールド
Русский


```

NOTE: debug flag will append itself to the text that is translated and I have no idea why. So don't use unless
there is no reponse
