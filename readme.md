Dead simple translator

```
go install gitlab.com/lenivyyluitel/granslater
```
or
```sh
git clone https://gitlab.com/lenivyyluitel/granslater
go build .
./granslater
```
granslater --help for all the commands

```

granslater --o ja hello world
ヘロワールド

                   >>error<<
granslater --o ru --s jp ヘロワールド --d true
map[error:jp is not supported]
<nil>

                   >>proper<<
./granslater --o ru --s ja ヘロワールド --d true
map[translatedText:Helloworld Торре]
Helloworld Торре


```
