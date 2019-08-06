
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
The webpage of api docs served at `http://localhost/api/docs/index.html`
<br>

![img](https://i.imgur.com/ds4k7tx.png)

## Database model
To fullfil all the API requirement, table `history` (from `varchar`, to `varchar`, rate `real`, date `date`) is defined. All the operation will be funneled into 1 table

## Test Case
These are workaround to cover all api test cases
- Add
```sh
$ curl -X POST "http://localhost/api/add/JMC/JPY" -H "accept: application/json"
```
- Remove
```sh
$ curl -X DELETE "http://localhost/api/remove/JMC/JPY" -H "accept: application/json"
```
- Track
```sh
$ curl -X GET "http://localhost/api/track?date=2018-11-13" -H "accept: application/json"
```
- Trend
```sh
$ curl -X GET "http://localhost/api/trend/USD/EUR?avg=1.0&vrn=0.1" -H "accept: application/json"
```
- Input
```sh
$ curl -X POST "http://localhost/api/input/AUD/USD?rate=0.73&date=2018-11-16" -H "accept: 
application/json"
```
- Remove
```sh
$ curl -X DELETE "http://localhost/api/remove/AUD/USD" -H "accept: application/json"
```
These api test have been covered up in test package along with db model

## Note
- docker-compose can't really control when to boot the app container, thats why I put restart as always in app service so the connection to db always retried until database ready to accept connection
- Before running the test package, the database and the environment variable should be well defined
- The app image hosted in google container registry, but the multistaged-build still can be observed in Dockerfile
- I use packr to bundle the static files so it compiled into binary in build time
- I use swagger (swaggo) annotation for documenting all functionality
