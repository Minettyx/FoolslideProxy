const htmlParser = require('node-html-parser');
const got = require('got');

var Provider = (function() {
    let methods = {};
  
    methods.search = async function(query) {
        if(query.length < 2) {
            return [];
        }

        var page1 = await got('https://juinjutsureader.ovh/directory/1/');
        const parsed = htmlParser.parse(page1.body);

        var pages_count = parsed.querySelector(".gbuttonright.btn.btn-primary").getAttribute("href").split("directory/")[1].split("/")[0];
        
        var pages = [];
        pages.push(page1.body);
        for(var i=2; i<=pages_count; i++) {
            var p = await got('https://juinjutsureader.ovh/directory/'+i+'/');
            pages.push(p.body);
        }

        var chs = [];
        pages.forEach(el => {
            var $ = htmlParser.parse(el);
            var caps = $.querySelectorAll(".series_element");
            caps.forEach(capEl => {
                var cap = {};
                cap.title = capEl.querySelector(".title").querySelector("a").text;
                cap.url = "/series/jj¥"+capEl.querySelector(".title").querySelector("a").getAttribute("href").split("series/")[1];
                chs.push(cap);
            });
        });

        var results = [];
        chs.forEach(ele => {
            if(ele.title.toLowerCase().includes(query.toLowerCase())) {
                results.push(ele);
            }
        });

        return results;
    }

    methods.manga = async function(id) {
        var page = await got('https://juinjutsureader.ovh/series/'+id);
        const $ = htmlParser.parse(page.body);

        var data = {};
        data.img = $.querySelector("img.thumb").getAttribute("src");
        data.synopsis = $.querySelector(".trama") ? $.querySelector(".trama").text.substring(7) : "";
        data.author = $.querySelector(".autore") ? $.querySelector(".autore").text.substring(8) : ""; 
        data.artist = $.querySelector(".artista") ? $.querySelector(".artista").text.substring(9) : ""; 

        var rawvolumes = $.querySelectorAll(".group_comic");
        var volumes = [];
        rawvolumes.forEach(el => {
            volume = {};
            volume.index = el.querySelector(".volume_comic").text.includes("VOLUME") ? parseInt(el.querySelector(".volume_comic").text.includes("VOLUME").substring(7)) : 0;

            var rawchs = el.querySelectorAll(".element");

            volume.chapters = [];
            rawchs.forEach(rawch => {
                var ch = {}
                ch.title = rawch.querySelector("a").text;
                ch.url = "/read/jj¥"+rawch.querySelector("a").getAttribute("href").split("read/")[1].split('/').join('ƒ');
                ch.date = psDate(rawch.querySelector(".meta_r").text);
                volume.chapters.push(ch);
            });

            volumes.push(volume);
        });

        data.volumes = volumes;


        return data;
    }

    methods.chapter = async function(chid) {
        var page = await got('https://juinjutsureader.ovh/read/'+chid);
        return JSON.parse(page.body.split("var pages = ")[1].split(";")[0]);
    }
  
    return methods;
})();

function psDate(date) {
    switch (date) {
        case "Oggi":
            var now = new Date();

            return now.getFullYear()+"."+("0"+(now.getMonth()+1)).slice(-2)+"."+now.getDate();
            break;
        case "Ieri":
            var now = new Date();
            return now.getFullYear()+"."+("0"+(now.getMonth()+1)).slice(-2)+"."+(now.getDate()-1);
            break;
        default:
            return date;
    }
}

module.exports = Provider;