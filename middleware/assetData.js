var jwt = require('json-web-token');
// var payload = {
//     "iss": "my_issurer",
//     "aud": "World",
//     "iat": 1400062400223,
//     "typ": "/online/transactionstatus/v2",
//     "request": {
//         "myTransactionId": "[myTransactionId]",
//         "merchantTransactionId": "[merchantTransactionId]",
//         "status": "SUCCESS"
//     }
// };
var secret = 'a2werrasd23QWdsfqwe23';
// encode
jwt.encode(secret, payload, function(err, token) {
    if (err) {
        console.error(err.name, err.message);
    } else {
        console.log(token);
        // decode
        jwt.decode(secret, token, function(err, decodedPayload,
            decodedHeader) {
            if (err) {
                console.error(err.name, err.message);
            } else {
                console.log(decodedPayload, decodedHeader);
            }
        });
    }
});