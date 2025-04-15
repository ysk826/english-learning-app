FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm install

# Copy the rest of the application
COPY . .

# Expose port
EXPOSE 3000

# Start application with host set to 0.0.0.0 to make it accessible outside the container
ENV HOST=0.0.0.0
CMD ["npm", "start"]