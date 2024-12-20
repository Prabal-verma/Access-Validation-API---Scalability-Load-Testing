# Game Authorization Service

A high-performance authorization service that validates rule-based access control for different games. The service supports dynamic rules based on geo-location, app version, platform, and app type.

## Features

### Core Features
- Rule-based access control
- Real-time validation
- High-performance (1M requests/minute)
- Redis caching
- Admin dashboard for rule management
- Test console for validation testing
- Performance metrics visualization

### Technical Specifications
- Backend: Go
- Frontend: React
- Database: Redis
- API: RESTful

## System Architecture

```
┌─────────────┐     ┌──────────────┐     ┌─────────────┐
│   React     │     │     Go       │     │    Redis    │
│  Frontend   │────▶│   Backend    │────▶│   Cache     │
└─────────────┘     └──────────────┘     └─────────────┘
     UI Layer         Service Layer        Data Layer
```

## API Documentation

### Validate Access

**Endpoint:** `POST /validate`

**Request:**
```json
{
    "game_id": "game1",
    "country": "US",
    "app_version": "1.0.0",
    "platform": "iOS",
    "app_type": "mobile"
}
```

**Response:**
```json
{
    "allowed": true
}
```

or

```json
{
    "allowed": false,
    "reason": "country not allowed"
}
```

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- Node.js and npm
- Docker

### Backend Setup

1. Clone the repository
```bash
git clone <repository-url>
cd <repository-name>
```

2. Start Redis
```bash
docker run -d -p 6379:6379 --name redis redis:7-alpine
```

3. Build and run the Go service
```bash
go build
./auth-service
```

### Frontend Setup

1. Navigate to the frontend directory
```bash
cd auth-service-dashboard
```

2. Install dependencies
```bash
npm install
```

3. Start the development server
```bash
npm start
```

## Usage Guide

### Managing Rules
1. Access the admin dashboard at `http://localhost:3000/admin`
2. Use the form to add new rules
3. View and manage existing rules in the table

### Testing Access
1. Go to the test console at `http://localhost:3000/test`
2. Fill in the request parameters
3. Click "Test Access" to see the validation result

### Monitoring Performance
1. View metrics at `http://localhost:3000/metrics`
2. Monitor:
   - Requests per second
   - Average latency
   - Error rate

## Performance

The service is designed to handle:
- 1 million requests per minute
- Average response time < 10ms
- Concurrent validations

### Scaling Strategies
1. Redis cluster for distributed caching
2. Load balancing for the Go service
3. React app deployment on CDN

## Security Considerations

1. All rules are stored securely in Redis
2. Backend validation for all requests
3. Rate limiting on API endpoints
4. CORS configuration for frontend


