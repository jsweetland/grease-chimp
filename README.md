# Grease Chimp

## Getting Started

### Dependencies

- Go
- NodeJS
- Docker

## Step 1 - Database

### Build the Database Image

If you have never started the database before, you will need to build the Docker image by running the following command.

```bash
./db/build.sh
```

### Start the Database

```bash
./db/start.sh
```

## Start the Backend

```bash
./api/start.sh
```

## Start the Frontend

```bash
./ui/start.sh
```
