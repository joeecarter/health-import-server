# Health Import Server
Storage backend for https://www.healthexportapp.com

Currently just stores the metrics into influxdb but more storage backends (and storing workout data) may be supported in the future.

## Config file
You'll need provide a json config file with the details on how to connect to and authenticate with your influx db instance:
```
[
	{
		"type": "influxdb",
		"hostname": "YOUR HOSTNAME HERE",
		"token": "YOUR TOKEN HERE",
		"org": "YOUR ORG HERE",
		"bucket": "YOUR BUCKET HERE"
	}
]
```

## Running in docker
The image can be built with this command (not on dockerhub yet):
```
docker build -t health-import:latest
```

To provide the config file to the application you need to place it here: /config/config.json (later on I'd like to support config via environment variables instead).

You can either do this with a bind mount e.g.
```
docker run -v $(PWD)/config:/config health-import:latest
```

Or making an image which extends the base image:
```
FROM docker-registry.home/joe/health-import:latest
ADD config.json /config/config.json
```
(docker-compose works well with this approach)

## What the metrics look like
See this file: [sample.go](/request/sample.go)

## How to use this with Health Export App (aka. API Export)
1. Run the server on a machine on your local home network.
2. Configure the API Export to point to the server.
3. Enable automatic syncing 
