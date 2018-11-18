# restful-forex
restful, tested, fully documented, container based golang application

## quick start
```sh
$ docker-compose up
```
This will spin docker compose chain tooling, to build, pull, and run necessary image as container. When this log appear for the last time, the application is ready to go.
> LOG:  database system is ready to accept connections

## web page
The webpage served at `http://localhost/web`
- Landing page
![img](https://imgur.com/g1ZwAbw)
- Input page
![img](https://imgur.com/mkWqN0L)
- Track page
![img](https://imgur.com/MqH9ld6)
- Trend page
![img](https://imgur.com/BVngeQb)
- Add page
![img](https://imgur.com/xgIFPzf)
- Remove page
![img](https://imgur.com/MZcp43O)

## test case
These are workaround to cover all api test cases
- add
```sh
$ curl -X POST "http://localhost:8080/api/add/JMC/JPY" -H "accept: application/json"
```
- remove
```sh
$ curl -X DELETE "http://localhost:8080/api/remove/JMC/JPY" -H "accept: application/json"
```
- track
```sh
$ curl -X GET "http://localhost:8080/api/track?date=2018-11-13" -H "accept: application/json"
```
- trend
```sh
$ curl -X GET "http://localhost:8080/api/trend/USD/EUR?avg=1.0&vrn=0.1" -H "accept: application/json"
```
- input
```sh
$ curl -X POST "http://localhost:8080/api/input/AUD/USD?rate=0.73&date=2018-11-16" -H "accept: 
application/json"
```
- remove
```sh
$ curl -X DELETE "http://localhost:8080/api/remove/AUD/USD" -H "accept: application/json"
```

## api docs
The webpage of api docs served at `http://localhost/api/docs`
![img](https://imgur.com/ds4k7tx)
