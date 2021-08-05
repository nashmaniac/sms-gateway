# SMS Gateway Api Documentation - V1

## Send SMS Api

**URL:** https://dev-sms-gateway.qikcheck.com/sms/v1/send

**Query Params**

|Params|DataType|Required|Description|Default Value
|--|--|--|--|--|
|apiKey|string|Yes|This is the api key that would be provided by SmartLab|
|to|string|Yes|Phone number where the sms would be sent|
|source|string|No|Language type of message|en
|message|string|Yes|Message that you want to send to receiver|
|override|boolean|No|true, if you want to send your custom message otherwise false|false


**Request Sample**

Here is the cURL snippet for the sms send api.
```sh
curl --location --request GET 'https://dev-sms-gateway.qikcheck.com/sms/v1/send?apiKey=<your_api_key>&to=<your_receiver_number>&source=en&message=<message_you_want_to_send>&override=false'
```
