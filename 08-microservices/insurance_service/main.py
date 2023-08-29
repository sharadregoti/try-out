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
ACCOUNT_SERVICE_URL = os.environ.get("INSURANCE_SERVICE_URL", "http://localhost:5000")

@app.route('/buyInsurance', methods=['POST'])
def buy_insurance():
    logging.info('Received request to buy insurance.')
    return jsonify({'message': 'Insurance purchase request received!'}), 200

@app.route('/getPolicyDetails/<policyId>', methods=['GET'])
def get_policy_details(policyId):
    logging.info(f'Received request to get details for policy ID: {policyId}.')
    return jsonify({'message': f'Policy details request received for policy ID: {policyId}!'}), 200

@app.route('/claimInsurance', methods=['POST'])
def claim_insurance():
    logging.info('Received request to claim insurance.')
    return jsonify({'message': 'Insurance claim request received!'}), 200

@app.route('/getClaimStatus/<claimId>', methods=['GET'])
def get_claim_status(claimId):
    logging.info(f'Received request to get claim status for claim ID: {claimId}.')
    return jsonify({'message': f'Claim status request received for claim ID: {claimId}!'}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
