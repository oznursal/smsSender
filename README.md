# smsSender
Scheduler Sms Sender project is  an automatic message sending system. 

    In this project, I am tasked with developing a system that automatically sends 2
    messages retrieved from the database(postgredb), which have not yet been sent, every 2 minutes, as
    illustrated in the CURL request provided below.

## Clone the project

```
$ git clone https://github.com/oznursal/smsSender.git
$ cd smsSender
```

## Rest-Api Curls

### Get Sent Messages


##### /api/v1/messages

    curl --location 'http://localhost:8080/api/v1/messages' \
    --header 'Content-Type: application/json'

###### Example response 

    [
    {
        "id": 5,
        "content": "   ",
        "phoneNumber": "905555555555",
        "sendingStatus": true
    },
    {
        "id": 6,
        "content": "232",
        "phoneNumber": "905555555555",
        "sendingStatus": true
    }
    ]


##### /api/v1/scheduler/start

    curl --location --request PUT 'http://localhost:8080/api/v1/scheduler/start'


##### /api/v1/scheduler/start

    curl --location --request PUT 'http://localhost:8080/api/v1/scheduler/stop'

## Postman Collection

   [sms_sender.postman_collection](docs/sms_sender.postman_collection)

## Webhook url
    https://webhook.site/#!/view/40850de7-286a-447e-abba-a760d20326cb

## How to work
    docker build --tag sms-sender .

    docker run -d --restart=always -p 8080:8080 sms-sender

### Docker-compose 

    docker compose up

Topics covered:

* Command-line flags ([flag](//pkg.go.dev/flag/))
* Web servers ([net/http](//pkg.go.dev/net/http/))
* Time ([time](//pkg.go.dev/time/))
* Logging ([log](//pkg.go.dev/log/))
* Postgre ([pg](//pkg.go.dev/github.com/go-pg/pg/))
* Redis ([redis](//pkg.go.dev/github.com/go-redis/redis/))
