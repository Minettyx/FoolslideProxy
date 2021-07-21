import { SearchResult, Manga, Chapter } from '../classes/interfaces'
import Module from '../classes/Module'
import axios from 'axios'
import cheerio from 'cheerio'

class Mangaworld implements Module {
  id = 'mw'
  name = 'MangaWorld'

  search(query: string) {
    return new Promise(async (resolve: (value: SearchResult[]) => void) => {

      const $ = cheerio.load((await axios.get(`https://www.mangaworld.io/archive?keyword=${query}`)).data)

      let results: SearchResult[] = []
      $('.entry').each((i, el) => {

        const mid = $('a', el).attr('href')?.split('/manga/')[1]+''
        const title = $('a', el).attr('title')+''

        results.push({id: mid, title: title})
      })

      resolve(results)

    })
  }

  manga(mangaid: string) {
    return new Promise(async (resolve: (value: Manga) => void) => {

      let chapters: Chapter[] = []

      let data = {
        synopsis: '',
        author: '',
        artist: '',
        img: '',
        chapters: chapters,
        sourceurl: 'https://www.mangaworld.io/manga/'+mangaid
      }

      const $ = cheerio.load((await axios.get(`https://www.mangaworld.io/manga/${mangaid}`)).data)

      data.synopsis = $('#noidungm').text()

      $(".col-12.col-md-6", $('.meta-data')).each((i, div) => {
        if($("span", div).text().includes('Autor')) {
          data.author = $("a", div).text()
        }
        if($("span", div).text().includes("Artist")) {
          data.artist = $("a", div).text()
        }
      })

      data.img = $("img.rounded").attr("src")+'';

      $('.chapter').each((i, entry) => {
        let cap = {
          title: '',
          id: '',
          date: new Date()
        }

        const chap = $('.chap', entry)

        cap.title = $('span', chap).text()
        cap.id = chap.attr('href')?.split('/read/')[1]+''
        cap.date = psDate($('i', chap).text())

        data.chapters.push(cap)
      })

      resolve(data)
    })
  }

  chapter(manga: string, id: string) {
    return new Promise(async (resolve: (value: string[]) => void) => {

      let result: string[] = []
      
      const body = (await axios.get(`https://www.mangaworld.io/manga/${manga}/read/${id}`)).data
      const json = JSON.parse(body.split('$MC=(window.$MC||[]).concat(')[1].split(')</script>')[0]);
      const pages: string[] = json.o.w[0][2].chapter.pages
      const $ = cheerio.load(body)

      const firstimage = $('#page').find('img').attr('src')?.split('/')
      firstimage?.pop()
      const baseurl = firstimage?.join('/')

      pages.forEach(page => {
        result.push(baseurl+'/'+page)
      })

      resolve(result)

    })
  }
}

export default new Mangaworld()

function psDate(input: string) {
  let mese = 0;
  let parts = input.split(' ');

  switch (parts[1]) {
      case "Gennaio":
          mese = 0;
          break;
  
      case "Febbraio":
          mese = 1;
          break;

      case "Marzo":
          mese = 2;
          break;

      case "Aprile":
          mese = 3;
          break;

      case "Maggio":
          mese = 4;
          break;
  
      case "Giugno":
          mese = 5;
          break;

      case "Luglio":
          mese = 6;
          break;

      case "Agosto":
          mese = 7;
          break;

      case "Settembre":
          mese = 8;
          break;
  
      case "Ottoble":
          mese = 9;
          break;

      case "Novembre":
          mese = 10;
          break;

      case "Dicembre":
          mese = 11;
          break;
  }

  let date = new Date()
  date.setFullYear(parseInt(parts[2]))
  date.setMonth(mese)
  date.setDate(parseInt(parts[0]))

  return date
}