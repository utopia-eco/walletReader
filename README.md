# walletReader

Request:
```json
# json
{
    “wallet_address”: string,
    “tokens”: [
        {“token_address”: string,
        “symbol”: string *omitempty
        “amount”: float64}, …
        ] *omitempty
}
```

Response:
```json
# json
{
    “wallet_address”: string,
    “wallet_value”: float64,
    “token_values”: [
        {“token_address”: string,
        “unit_price”: float64,
        “token_value”: float64}, …
        ]
}
```