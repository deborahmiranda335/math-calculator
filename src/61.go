var express = require('express');
var app = express();

app.get('/', function(req, res) {
  var numbers = [2, 4, 6, 8];
  var result = numbers.reduce(function(acc, curr) {
    return acc + curr * (curr - 1);
  }, 0);

  res.send('The sum of squares from 0 to ' + numbers.length + ' is: ' + result);
});

app.listen(3000);
