const express = require('express');
const router = express.Router();
const mangaworld = require('../sources/mangaworld');

router.post('/', async (req, res, next) => {try {

  res.set('Content-Type', 'text/html');

  var data = [];

  if(req.body.search) {
    data = data.concat(await mangaworld.search(req.body.search));
  }

  var response = "";
  data.forEach(ele => {
    response += `<div class="group"><div class="title"><a href="${ele.url}" title="${ele.title}">${ele.title}</a></div></div>`;
  });
  res.send(response);

} catch(err){next(err); console.log(err); }});

module.exports = router;
