const request = require('request');
const apiKey = 'bfd0bdb4833fb6d36e047a68fc95bbbd'
const cityName = 'London'
const url = 'api.openweathermap.org/data/2.5/weather?id='+cityName+'&appid='+apiKey


request(url, { json: true }, (err, res, body) => {
  if (err) { return console.log(err); }
  console.log(body.url);
  console.log(body.explanation);
});