# Deployment

This guide provides step-by-step instructions for deploying the Stadia application, covering both backend and frontend deployment options. Whether you prefer using Docker for a streamlined deployment or traditional server methods, this guide has you covered.

## Backend Deployment

- **Option 1**: Docker (Recommended)

```bash
# Backend directory contains a Dockerfile for easy containerization
# Build the Docker image
docker build -t stadia-backend .

# Run the container
docker run -p 8000:8000 stadia-backend
```

- **Option 2**: Traditional Server

```bash
# Build the binary
go build -o stadia-backend main.go

# Run the server
./stadia-backend
```

---

### Frontend Deployment

- **Build for Production**

```bash
cd frontend
npm run build
```

> This creates an optimized production build in the `dist/` directory.

#### Deploy to Netlify

```bash
# Install Netlify CLI
npm install -g netlify-cli

# Deploy
netlify deploy --prod --dir=dist
```

#### Deploy to Vercel

```bash
# Install Vercel CLI
npm install -g vercel

# Deploy
vercel --prod
```

### Environment Variables

**Backend** (Optional):

```bash
MODE=prod
PORT=8000
GIN_MODE=release
VERSION=1.0.0
NAME=Stadia
```

**Frontend**:

```env
VITE_API_URL=https://your-backend-url.com/api
```
