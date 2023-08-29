from opentelemetry.instrumentation.requests import RequestsInstrumentor
from flask import Flask, request, jsonify
import os
import requests
import logging

RequestsInstrumentor().instrument()

app = Flask(__name__)

# Initialize logging
logging.basicConfig(level=logging.INFO)

# Retrieve the dependent service URL from an environment variable
ACCOUNT_SERVICE_URL = os.environ.get("LOAN_SERVICE_URL", "http://localhost:5000")
INSURANCE_SERVICE_URL = os.environ.get("INSURANCE_SERVICE_URL", "http://localhost:5000")

@app.route('/applyLoan', methods=['POST'])
def apply_loan():
    logging.info('Received request to apply for a loan.')
    return jsonify({'message': 'Loan application request received!'}), 200

@app.route('/getLoanStatus/<loanId>', methods=['GET'])
def get_loan_status(loanId):
    logging.info(f'Received request to get status for loan ID: {loanId}.')
    return jsonify({'message': f'Loan status request received for loan ID: {loanId}!'}), 200

@app.route('/getLoanDetails/<loanId>', methods=['GET'])
def get_loan_details(loanId):
    logging.info(f'Received request to get details for loan ID: {loanId}.')
    return jsonify({'message': f'Loan details request received for loan ID: {loanId}!'}), 200

@app.route('/calculateEMI', methods=['POST'])
def calculate_emi():
    logging.info('Received request to calculate EMI.')
    # Making a call to InsuranceService's getPolicyDetails endpoint
    response = requests.get(f"{INSURANCE_SERVICE_URL}/getPolicyDetails/12345")  # Using a dummy policy ID for demonstration
    logging.info(f'Received response from InsuranceService: {response.text}')
    return jsonify({'message': 'EMI calculation request received!'}), 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
