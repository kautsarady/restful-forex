# restful-forex
restful, tested, fully documented, container based golang application

## Quick Start
```sh
$ docker-compose up
```
This will spin docker compose chain tooling, to build, pull, and run necessary image as container. When this log appear for the last time, the application is ready to go.
> LOG:  database system is ready to accept connections

## Web Page
The webpage served at `http://localhost/web`
- Landing Page
<br>
![img](https://i.imgur.com/g1ZwAbw.png)

- Input Page
<br>
![img](https://i.imgur.com/mkWqN0L.png)

- Track Page
<br>
![img](https://i.imgur.com/MqH9ld6.png)

- Trend Page
<br>
![img](https://i.imgur.com/BVngeQb.png)

- Add Page
<br>
![img](https://i.imgur.com/xgIFPzf.png)

- Remove Page
<br>
![img](https://i.imgur.com/MZcp43O.png)

## API docs
The webpage of api docs served at `http://localhost/api/docs`
<br>
![img](https://i.imgur.com/ds4k7tx.png)


## Test Case
These are workaround to cover all api test cases
- Add
```sh
$ curl -X POST "http://localhost:8080/api/add/JMC/JPY" -H "accept: application/json"
```
- Remove
```sh
$ curl -X DELETE "http://localhost:8080/api/remove/JMC/JPY" -H "accept: application/json"
```
- Track
```sh
$ curl -X GET "http://localhost:8080/api/track?date=2018-11-13" -H "accept: application/json"
```
- Trend
```sh
$ curl -X GET "http://localhost:8080/api/trend/USD/EUR?avg=1.0&vrn=0.1" -H "accept: application/json"
```
- Input
```sh
$ curl -X POST "http://localhost:8080/api/input/AUD/USD?rate=0.73&date=2018-11-16" -H "accept: 
application/json"
```
- Remove
```sh
$ curl -X DELETE "http://localhost:8080/api/remove/AUD/USD" -H "accept: application/json"
```