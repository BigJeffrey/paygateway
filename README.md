# paygateway
To run this you need to set environment variable:
PGP=qwerty12345678!

and set postgres database from file in doc folder (paygatewaydump) or use provided one (It is docker container on private VPS).

Test are provided by attached file in doc folder - "PayGateway.postman_collection.json" or https://documenter.getpostman.com/view/16998762/UVeMKQFY

You need to use valid credit card numbers because of Luhn check implemented.
To test specific credit card numbers that was provided by the task they are checked before Luhn verification cause they would not pass it.
