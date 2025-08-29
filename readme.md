## Web scrapper cli

#### Made with go(goquery, cobra)

#### Clean architecture

### Installing

Firstly add .env file, I recommend don't change it
`cp .env.example .env`

### Building

`sudo docker build -t site-scrapper .`

### Executing

You can remove -d flag if you want
`sudo docker run -d -it -p site-scrapper`

### Using

Secondly go inside docker container by using command
`sudo docker exec -it ${site-scrapper id of container} sh`
There will be executable file scrapper
It is your cli, you can use it

### Example of View

By going to
`https://localhost:8008/example.ya.ru.html`
You will see example of already executed command
`./scrapper --url https://google.com --keyword yandex`
