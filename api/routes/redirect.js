const express = require('express');
const router = express.Router();
const apicache = require('apicache');

var cache = apicache.middleware;

router.get('/:id', cache('5 minutes'), async (req, res, next) => {try {

  res.redirect(decodeURIComponent(req.params.id));

} catch(err){next(err);console.log(err);}});

module.exports = router;
