import express from 'express'
import morgan from 'morgan'
import bodyParser from 'body-parser'
const app = express()

import mainRoute from './main'

app.use(morgan('dev'))
app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

app.use((req, res, next) => {
  res.set('Access-Control-Allow-Origin', '*')
  res.set('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept, Authorization')
  if(req.method === 'OPTIONS') {
    res.set('Access-Control-Allow-Methods', 'PUT, POST, PATCH, DELETE, GET')
    return res.status(200).json({})
  }
  next()
})

app.use('/', mainRoute);

//ERROR handling
app.use((req, res, next) => {
  res.status(404).json({message: 'Cannot '+req.method+' '+req.path, status: 404, name: "NotFound"})
})

app.listen(process.env.PORT || 3000, () => console.log("Listening"))