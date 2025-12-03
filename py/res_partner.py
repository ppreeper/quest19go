#!/usr/bin/env python3
# You MUST store this key securely. Place it in an
# environment variable or in in a file outside of
# git (e.g. your home directory).
api_key = "5ef6cf19b1c270ad4489095e9839b80aec61d32e"

import requests
import json
response = requests.post(
    "https://quest19data.odoopro.ca/json/2/product.template/search_read",
    headers={
        "Authorization": f"Bearer {api_key}",
        # "X-Odoo-Database": "...",
    },
    json={
        "domain": [
            [
                "name",
                "ilike",
                "a%"
            ]
        ],
        "fields": [
            "name"
        ],
        "limit": 20
    },
)
response.raise_for_status()
data = response.json()
print(json.dumps(data, indent=4))


response = requests.post(
    "https://quest19data.odoopro.ca/json/2/res.partner/search",
    headers={
        "Authorization": f"Bearer {api_key}",
        # "X-Odoo-Database": "...",
    },
    json={
        "domain": [
            [
                "name",
                "ilike",
                "a%"
            ]
        ]
    },
)
response.raise_for_status()
data = response.json()
print(json.dumps(data, indent=4))


# You MUST store this key securely. Place it in an
# environment variable or in in a file outside of
# git (e.g. your home directory).
response = requests.post(
    "https://quest19data.odoopro.ca/json/2/res.partner/create",
    headers={
        "Authorization": f"Bearer {api_key}",
        # "X-Odoo-Database": "...",
    },
    json={
        "vals_list": [
            {
                "name": "Peter Preeper",
                "company_type": "company",
            }
        ]
    },
)
response.raise_for_status()
data = response.json()
print(json.dumps(data, indent=4))
