const express = require('express');
const router = express.Router();
const apicache = require('apicache');
const mangaworld = require('../sources/mangaworld');

var cache = apicache.middleware;

router.post('/:manga/:volume/:chapter', cache("5 minutes"), async (req, res, next) => {try {

  res.set('Content-Type', 'text/html');

  var data = [];

  if(req.params.manga.includes("mw¥")) {
    data = await mangaworld.chapter(req.params.manga.split("¥")[1].replace('ƒ', '/'), req.params.chapter);
  } else {
    res.status(404).json({});
    return;
  }

  res.send("var pages = "+JSON.stringify(data)+";");

} catch(err){next(err); console.log(err); }});

module.exports = router;
