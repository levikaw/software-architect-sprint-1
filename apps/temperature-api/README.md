# Temperature API

## Prerequisites

- Docker

## Getting Started

### Option 1: Using Docker (Recommended)

The easiest way to start the application is to use Docker:


```bash
docker build -t temperatureapi:latest .
docker run -d -p 8081:8081 temperatureapi:latest 
```

The API will be available at [http://localhost:8081](http://localhost:8081)

### Option 2: Native run (Golang required)

If you prefer to run the application without Docker, run following command from `src` dir:

```bash
go run *.go
```

## API Endpoints

- `GET /temperature` - return int value of temperature, allow `location` and `sensorId` parameters

