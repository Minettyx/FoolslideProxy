const express = require('express');
const router = express.Router();
const apicache = require('apicache');
const mangaworld = require('../sources/mangaworld');
const juinjutsu = require('../sources/juinjutsu');

var cache = apicache.middleware;

router.post('/:chapter', cache("5 minutes"), async (req, res, next) => {try {

  res.set('Content-Type', 'text/html');

  var data = [];

  if(req.params.chapter.includes("mw¥")) {
    data = await mangaworld.chapter(req.params.chapter.split("¥")[1].split('ƒ').join('/'));
  } else if(req.params.chapter.includes("jj¥")) {
    data = await juinjutsu.chapter(req.params.chapter.split("¥")[1].split('ƒ').join('/'));
  } else {
    res.status(404).json({});
    return;
  }

  res.send("var pages = "+JSON.stringify(data)+";");

} catch(err){next(err); console.log(err); }});

module.exports = router;
