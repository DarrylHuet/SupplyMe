const axios = require('axios');

axios.get('/user?ID=12345')
	.then(function (response){
		console.log(response);
	})
	.catch(function (error) {
		console.log(error);
	})
	.then(function () {

	});

asynx function getUser() {
	try {
		const reponse = await axios.get('user?ID=12345');
		console.log(response);
	} catch (error) {
		console.error(error);
	}
}
