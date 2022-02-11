# space_trouble



# Curl tests

## Create Booking
curl -X POST http://127.0.0.1:8080/booking -H 'Content-Type: application/json' -d '{
    "first_name":"name",
    "last_name":"last",
    "gender": "gender",
    "birthday": "birthday",
    "launchpad_id": "launchpad_id",
    "destination_id": "destination_id",
    "launch_date": "launch_date"
}'


## Get Bookings
curl -X GET http://127.0.0.1:8080/booking


## Delete Booking
curl -X DELETE http://127.0.0.1:8080/booking/ -H 'Content-Type: application/json' -d '{
    "booking_id":"123"
}'