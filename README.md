# Roast API

A simple Go REST API that serves random roasts from a PostgreSQL database.

## Features

- Single endpoint `/api/roasts` that returns a random roast
- PostgreSQL database integration with connection pooling
- CORS enabled for cross-origin requests
- Environment variable configuration
- Deployed on Fly.io

## API Endpoints

### GET /api/roasts

Returns a random roast from the database.

**Response:**
```json
{
  "roast": "Your randomly selected roast text here"
}
```

**Status Codes:**
- `200` - Success
- `404` - No roasts found in database
- `405` - Method not allowed (only GET supported)
- `500` - Internal server error

## Environment Variables

Create a `.env` file with the following variables:

```env
PORT=8080
DATABASE_URL=postgresql://username:password@host:port/database
```

## Running Locally

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up your PostgreSQL database with a `roasts` table:
   ```sql
   CREATE TABLE roasts (
       id SERIAL PRIMARY KEY,
       content TEXT NOT NULL
   );
   ```
4. Create a `.env` file with your database configuration
5. Run the application:
   ```bash
   go run .
   ```

The API will be available at `http://localhost:8080`

## Building

To build the binary:
```bash
go build -o roast-api .
```

## Deployment

This project is configured for deployment on [Fly.io](https://fly.io). The configuration is in `fly.toml`.

To deploy:
```bash
flyctl deploy
```

## Technology Stack

- **Go 1.25.1** - Programming language
- **PostgreSQL** - Database
- **pgx/v5** - PostgreSQL driver with connection pooling
- **godotenv** - Environment variable loading
- **Fly.io** - Hosting platform

## Database Schema

The application expects a PostgreSQL table named `roasts`:

```sql
CREATE TABLE roasts (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL
);
```

Add your roasts to this table, and the API will randomly select one for each request.