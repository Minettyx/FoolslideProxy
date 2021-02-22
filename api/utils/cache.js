const NodeCache = require( "node-cache" );
const got = require('got');

const cache = new NodeCache( { stdTTL: 100, checkperiod: 120 } );

var CacheURL = (function() {
  let methods = {};

  methods.get = async function(id, callback, ttl = 60) {

    var data = cache.get(id);
    if(data == undefined) {
      cache.set(id, callback(), ttl);
    }

    return data;
  }

  return methods;
})();

module.exports = CacheURL;
