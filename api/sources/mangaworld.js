const htmlParser = require('node-html-parser');
const got = require('got');

var MangaCover = (function() {
    let methods = {};
  
    methods.search = async function(query) {

        var data = [];
        var page = await got('https://www.mangaworld.cc/archive?keyword='+query);
        const parsed = htmlParser.parse(page.body);

        var entrys = parsed.querySelectorAll('.entry');

        entrys.forEach(entry => {
            var a = entry.querySelector("a");
            var d = [];
            var mid = a.getAttribute("href").split('/manga/')[1].includes('/') ? a.getAttribute("href").split('/manga/')[1].split('/')[0] : a.getAttribute("href").split('/manga/')[1]
            d.url = '/series/mw¥'+mid
            d.title = a.getAttribute("title");
            data.push(d);
        });

        return data;
    }

    methods.manga = async function(id) {

        var data = [];
        var page = await got('https://www.mangaworld.cc/manga/'+id);
        const parsed = htmlParser.parse(page.body);

        data.synopsis = parsed.querySelector('#noidungm').innerHTML;

        parsed.querySelector('.meta-data').querySelectorAll(".col-12.col-md-6").forEach(div => {
            if(div.querySelector("span").innerHTML.includes('Autor')) {
                data.author = div.querySelector("a").innerHTML;
            }
            if(div.querySelector("span").innerHTML.includes("Artist")) {
                data.artist = div.querySelector("a").innerHTML;
            }
        });

        data.img = parsed.querySelector("img.rounded").getAttribute("src");

        data.volumes = [
            {
                index: 0,
                chapters: []
            }
        ];

        var entrys = parsed.querySelectorAll('.chapter');

        entrys.forEach(entry => {
            var chap = entry.querySelector('.chap');

            var cap = {}
            cap.title = chap.querySelector("span").innerHTML;
            cap.url = '/read/mw¥'+chap.getAttribute("href").split('/manga/')[1].split('/read/')[0].replace('/', "ƒ")+'/0/'+chap.getAttribute("href").split('/read/')[1];

            cap.date = psDate(chap.querySelector('i').innerHTML);

            data.volumes[0].chapters.push(cap);
        });

        return data;
    }

    methods.chapter = async function(manga, id) {

        var data = [];

        var page = await got('https://www.mangaworld.cc/manga/'+manga+'/read/'+id);

        var json = JSON.parse(page.body.split('$MC=(window.$MC||[]).concat(')[1].split(')</script>')[0]);

        //var pages = JSON.parse('['+page.body.split('"chapters":[')[1].split(']')[0].split('[')[1]+']');

        var pages = json.o.w[0][2].chapter.pages;

        const parsed = htmlParser.parse(page.body);

        var firstimage = parsed.querySelector("#page").querySelector('img').getAttribute("src").split('/');
        firstimage.pop();
        var baseurl = firstimage.join("/");

        pages.forEach(page => {
            var d = {};
            d.url = baseurl+"/"+page;
            data.push(d);
        });

        console.log(data);

        return data;
    }
  
    return methods;
})();
  
function psDate(input) {
    var mese = 0;
    var parts = input.split(' ');

    switch (parts[1]) {
        case "Gennaio":
            mese = 1;
            break;
    
        case "Febbraio":
            mese = 2;
            break;

        case "Marzo":
            mese = 3;
            break;

        case "Aprile":
            mese = 4;
            break;

        case "Maggio":
            mese = 5;
            break;
    
        case "Giugno":
            mese = 6;
            break;

        case "Luglio":
            mese = 7;
            break;

        case "Agosto":
            mese = 8;
            break;

        case "Settembre":
            mese = 9;
            break;
    
        case "Ottoble":
            mese = 10;
            break;

        case "Novembre":
            mese = 11;
            break;

        case "Dicembre":
            mese = 12;
            break;
    }

    return parts[2]+"."+mese+"."+parts[0];
}

module.exports = MangaCover;