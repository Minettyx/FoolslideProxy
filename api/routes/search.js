const express = require('express');
const router = express.Router();
const mangaworld = require('../sources/mangaworld');
const juinjutsu = require('../sources/juinjutsu');

router.post('/', async (req, res, next) => {try {

  res.set('Content-Type', 'text/html');

  var data = [];

  if(req.body.search) {
    if(req.body.search.toLowerCase().startsWith("mangaworld:")) {
      data = data.concat(await mangaworld.search(req.body.search.substring(11).trim()));
    } else if(req.body.search.toLowerCase().startsWith("mw:")) {
      data = data.concat(await mangaworld.search(req.body.search.substring(3).trim()));
      

    } else if(req.body.search.toLowerCase().startsWith("juinjutsu:")) {
      data = data.concat(await juinjutsu.search(req.body.search.substring(10).trim()));
    } else if(req.body.search.toLowerCase().startsWith("jj:")) {
      data = data.concat(await juinjutsu.search(req.body.search.substring(3).trim()));


    } else {
      data = data.concat(await mangaworld.search(req.body.search));
      data = data.concat(await juinjutsu.search(req.body.search));
    }
  }

  var response = "";
  data.forEach(ele => {
    response += `<div class="group"><div class="title"><a href="${ele.url}" title="${ele.title}">${ele.title}</a></div></div>`;
  });
  res.send(response);

} catch(err){next(err); console.log(err); }});

module.exports = router;
