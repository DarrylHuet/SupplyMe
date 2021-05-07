//  let myObj = {
//     name: "Chris",
//      age: 38
//  };
// myObj
const bcrypt = require('bcrypt');

let myString = JSON.stringify(myObj);
//Now Marshalling
request.open('GET', requestURL);
request.responseType = 'text'; // now we're getting a string!
request.send();
request.onload = function() {
    const stringToEncode = request.response; // get the string

    const stringEncoded = JSON.parse(stringToEncode); // convert it
    bcrypt.hash(stringEncoded, rounds, (err, hash) => {
        if (err) {
          console.error(err)
          return
        }
        console.log(hash)
    })
    
    fs.writeFile ("input.json", JSON.stringify(stringEncoded), function(err) {
    if (err) throw err;
    console.log('complete');
    }
);
  
}