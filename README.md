PORT=8080
DB_DRIVER=postgres
DB_HOST=fxpg
DB_PORT=5432
DB_DBNAME=forex
DB_USER=admin
DB_PASSWORD=kavedeveloper

# valid test
* add
- curl -X POST "http://localhost:8080/api/add/JMC/JPY" -H "accept: application/json"
* remove
- curl -X DELETE "http://localhost:8080/api/remove/JMC/JPY" -H "accept: application/json"
* track
- curl -X GET "http://localhost:8080/api/track?date=2018-11-13" -H "accept: application/json"
* trend
- curl -X GET "http://localhost:8080/api/trend/USD/EUR?avg=1.0&vrn=0.1" -H "accept: application/json"
* input
- curl -X POST "http://localhost:8080/api/input/AUD/USD?rate=0.73&date=2018-11-16" -H "accept: application/json"
* remove
- curl -X DELETE "http://localhost:8080/api/remove/AUD/USD" -H "accept: application/json"

TODO:
• provide all env var in dockerfile (not in .env), need refactor
• add always restart to forex app

docker run --name fxpg --hostname fxpg -p 5432:5432 -e POSTGRES_PASSWORD=kavedeveloper -e POSTGRES_DB=forex -e POSTGRES_USER=admin -d postgres