import hashlib
import hmac
import requests
import json
from requests.exceptions import HTTPError

# Replace with your actual API key and secret key
API_KEY = 'YOUR_API_KEY'
SECRET_KEY = 'YOUR_SECRET_KEY'

# Kucoin API base URL
BASE_URL = 'https://api.kucoin.com'

# Endpoint for getting account list
ENDPOINT = '/api/v1/accounts'

# Build the request URL
url = BASE_URL + ENDPOINT

# Set up the request headers
headers = {
    'Content-Type': 'application/json',
    'KC-API-KEY': API_KEY,
    'KC-API-SIGNATURE': '',
    'KC-API-TIMESTAMP': '',
    'KC-API-PASSPHRASE': '',
}

# Make sure to replace 'YOUR_SECRET_KEY' with your actual secret key
def generate_signature(api_key, secret_key, timestamp, method, endpoint, params):
    message = f'{timestamp}{method}{endpoint}{json.dumps(params)}'
    signature = hmac.new(secret_key.encode('utf-8'), message.encode('utf-8'), hashlib.sha256).hexdigest()
    return signature

# Get server timestamp
response = requests.get(BASE_URL + '/api/v1/timestamp')
timestamp = response.json()['data']

# Set up the request parameters
params = {
    'currency': '',  # You can specify a currency if needed
    'type': 'main',  # Account type: main, trade, margin, trade_hf
}

# Generate signature
signature = generate_signature(API_KEY, SECRET_KEY, timestamp, 'GET', ENDPOINT, params)

# Update headers with the generated signature and timestamp
headers['KC-API-SIGNATURE'] = signature
headers['KC-API-TIMESTAMP'] = timestamp

# Make the API request
try:
    response = requests.get(url, headers=headers, params=params)
    response.raise_for_status()  # Raise an HTTPError for bad responses
    data = response.json()
    print(json.dumps(data, indent=2))  # Print the response data in a readable format

except HTTPError as http_err:
    print(f'HTTP error occurred: {http_err}')
except Exception as err:
    print(f'An error occurred: {err}')
