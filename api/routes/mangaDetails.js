const express = require('express');
const router = express.Router();
const apicache = require('apicache');
const mangaworld = require('../sources/mangaworld');
const juinjutsu = require('../sources/juinjutsu');

var cache = apicache.middleware;

router.all('/:id', cache("5 minutes"), async (req, res, next) => {try {

  res.set('Content-Type', 'text/html');

  var data = {}

  var id = req.params.id.split("¥")[1];
  if(req.params.id.startsWith("mw¥")) {
    data = await mangaworld.manga(id);
  } else if(req.params.id.startsWith("jj¥")) {
    data = await juinjutsu.manga(id);
  } else {
    res.status(404).json({});
    return;
  }

  var response = `<html>
  <head></head>
  <body>
      <div id="wrapper">
          <article id="content">
              <div class="panel">
                  <div class="comic info">
                      <div class="thumbnail">
                          <img src="${data.img}" />
                      </div>
      
                      <div class="large comic">
                          <h1 class="title"></h1>
                          <div class="info">
                              <b>Author</b>: ${data.author}<br>
                              <b>Artist</b>: ${data.artist}<br>
                              <b>Synopsis</b>: ${data.synopsis}
                          </div>
                      </div>
                  </div><div class="list">`;

    data.volumes.forEach(volume => {
        response += `<div class="group"><div class="title">Volume ${volume.index}</div>`;
            volume.chapters.forEach(chapter => {
                response += `<div class="element">
                    <div class="title"><a href="${chapter.url}" title="${chapter.title}">${chapter.title}</a></div>
                    <div class="meta_r">by <a href="" title="" ></a>, ${chapter.date}</div>
                </div>`;
            });
        response += `</div>`;
    });

    response += `</div></div></article></div></body></html>`;

  res.send(response);

} catch(err){next(err); console.log(err); }});

module.exports = router;
