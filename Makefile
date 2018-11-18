build:
	docker build --no-cache -t gcr.io/kautsarady-forex-gcr/forex_image .

push:
	gcloud docker -- push gcr.io/kautsarady-forex-gcr/forex_image