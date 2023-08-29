from opentelemetry.instrumentation.requests import RequestsInstrumentor
from flask import Flask, jsonify
import os
import requests
import logging

RequestsInstrumentor().instrument()

app = Flask(__name__)

# Initialize logging
logging.basicConfig(level=logging.INFO)

# Retrieve the dependent service URL from an environment variable
LOAN_SERVICE_URL = os.environ.get("LOAN_SERVICE_URL", "http://localhost:5000")
INSURANCE_SERVICE_URL = os.environ.get("INSURANCE_SERVICE_URL", "http://localhost:5000")

@app.route('/createAccount', methods=['POST'])
def create_account():
    logging.info('Received request to create account.')
    return jsonify({'message': 'Account creation request received!'}), 200

@app.route('/getBalance/<accountId>', methods=['GET'])
def get_balance(accountId):
    logging.info(f'Received request to get balance for account ID: {accountId}.')
    # Making a call to LoanService's calculateEMI endpoint
    response = requests.post(f"{LOAN_SERVICE_URL}/calculateEMI")
    logging.info(f'Received response from LoanService: {response.text}')
    return jsonify({'message': f'Balance request received for account ID: {accountId}!'}), 200

@app.route('/transferFunds', methods=['POST'])
def transfer_funds():
    logging.info('Received request to transfer funds.')
    return jsonify({'message': 'Fund transfer request received!'}), 200

@app.route('/accountHistory/<accountId>', methods=['GET'])
def account_history(accountId):
    logging.info(f'Received request for transaction history for account ID: {accountId}.')
    return jsonify({'message': f'Transaction history request received for account ID: {accountId}!'}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
