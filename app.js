const express = require('express');
const app = express();
const morgan = require('morgan');
const bodyParser = require('body-parser');

const searchRoute = require('./api/routes/search');
const mangaDetailsRoute = require('./api/routes/mangaDetails');
const readRoute = require('./api/routes/read');

app.use(morgan('dev'));
app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

//CORS and OPTIONS setup
app.use((req, res, next) => {
  res.set('Access-Control-Allow-Origin', '*');
  res.set('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept, Authorization');
  if(req.method === 'OPTIONS') {
    res.set('Access-Control-Allow-Methods', 'PUT, POST, PATCH, DELETE, GET');
    return res.status(200).json({});
  }
  next();
});

app.use('/search', searchRoute);
app.use('/series', mangaDetailsRoute);
app.use('/read', readRoute);


//ERROR handling
app.use((req, res, next) => {
  const error = new Error('Cannot '+req.method+' '+req.path);
  error.status = 404;
  error.name = "NotFound";
  next(error);
})

app.use((error, req, res, next) => {
  res.status(error.status || 500);
  res.json({
    "error": error.name,
    "message": error.message,
    "status": error.status
  });
})

module.exports = app;
