# Use an official Python runtime as the base image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the local requirements.txt to the container
COPY requirements.txt ./

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the local code to the container
COPY . .

# Define the command to run your app using CMD
CMD ["opentelemetry-instrument", "python3", "main.py"]
