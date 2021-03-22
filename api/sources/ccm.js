const htmlParser = require('node-html-parser');
const got = require('got');

var Provider = (function() {
    let methods = {};
  
    methods.search = async function(query) {
        if(query.length < 2) {
            return [];
        }

        var page = await got('https://api.ccmscans.in/mangas');
        var content = JSON.parse(page.body);

        var data = [];
        content.forEach(el => {
            if(el.title.toLowerCase().includes(query.toLowerCase())) {
                data.push({title: el.title, url: '/series/ccm¥'+el.id});
            }
        });

        return data;
    }

    methods.manga = async function(id) {
        var page = await got('https://api.ccmscans.in/manga/'+id);
        var content = JSON.parse(page.body);

        var data = {};
        data.synopsis = '';
        data.author = content.author;
        data.artist = content.artist;
        data.img = content.cover;

        data.volumes = [
            {
                index: 0,
                chapters: []
            }
        ];

        content.chapters.forEach(ch => {
            var d = {};
            d.title = (ch.volume!='' ? 'Vol.'+ch.volume+' ' : '')+'Ch.'+ch.chapter+(ch.title!='' ? ' - '+ch.title : '');
            d.url = '/read/ccm¥'+id+'ƒ'+ch.chapter;
            var dt = new Date(parseInt(ch.time)*1000);
            d.date = dt.getFullYear()+"."+("0"+(dt.getMonth()+1)).slice(-2)+"."+dt.getDate();
            data.volumes[0].chapters.push(d);
        });

        data.volumes[0].chapters = data.volumes[0].chapters.reverse();

        return data;
    }

    methods.chapter = async function(chid) {
        var page = await got('https://api.ccmscans.in/chapter/'+chid);
        var images = JSON.parse(page.body).images;
        
        var data = [];
        images.forEach(img => {
            data.push({url: img});
        });

        return data;
    }
  
    return methods;
})();

module.exports = Provider;