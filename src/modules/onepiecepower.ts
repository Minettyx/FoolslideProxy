import { SearchResult, Manga, Chapter } from '../classes/interfaces'
import Module from '../classes/Module'
import axios from 'axios'
import cheerio from 'cheerio'

class Juinjutsu implements Module {
  id = 'opp'
  name = 'One Piece Power'

  search(query: string) {
    return new Promise(async (resolve: (value: SearchResult[]) => void) => {
      const page = await axios.get('http://onepiecepower.info/manga/lista-manga')
      const parsed = cheerio.load(page.data)

      const res: SearchResult[] = []
      await Promise.all(parsed('body > table > tbody > tr > td').find('a').each((i, e) => {
        const el = cheerio(e)
        if(el.text().toLowerCase().includes(query.toLowerCase())) {
          res.push({
            id: el.attr('href')+'',
            title: el.text()
          })
        }
      }))

      resolve(res)
    })
  }

  manga(mangaid: string) {
    return new Promise(async (resolve: (value: Manga) => void) => {
      const page = await axios.get('http://onepiecepower.info/manga/'+mangaid)
      const parsed = cheerio.load(page.data)

      const chapters: Chapter[] = []
      await Promise.all(parsed('body > table > tbody > tr:last-child > td').find('a').each((i, e) => {
        const el = cheerio(e)
        chapters.push({
          title: el.text(),
          id: el.attr('href')+'',
          date: new Date(0)
        })
      }))

      let aut = '', syn = ''
      parsed('body > table > tbody > tr:nth-child(3) > td').find('span').each((i, e) => {
        const el = parsed(e)
        if(el.text() === 'Autore:') {
          aut = el.next('em').text()
        }
        if(el.text() === 'Descrizione:') {
          syn = el.next('em').text()
        }
      })

      const img = new URL(
        'images/cover.jpg',
        "http://onepiecepower.info/manga/"+mangaid
      ).href

      resolve({
        synopsis: syn,
        author: aut,
        artist: '',
        img: img,
        chapters: chapters.reverse(),
        sourceurl: 'http://onepiecepower.info/manga/'+mangaid
      })
    })
  }

  chapter(manga: string, id: string) {
    return new Promise(async (resolve: (value: string[]) => void) => {
      const url = new URL(id, 'http://onepiecepower.info/manga/'+manga).href
      const html = (await axios.get(url)).data

      let oversize = 15
      while (true) {
        if(await this.pageExist(html, manga, id, oversize)) {
          oversize *= 2
        } else {
          break
        }
      }

      let res: string[] = []
      for (let i = 1; i <= oversize; i++) {
        res.push(this.pageUrl(html, manga, id, i))
      }

      resolve(res)

    })
  }

  private pageUrl(html: string, manga: string, chapter: string, page: number) {
    let url = new URL(chapter, 'http://onepiecepower.info/manga/'+manga).href
    let link = ''
    let vol = html.split('var vol = "')[1].split('"')[0]
    let cap = chapter.split('/')[1]+'/'

    link=url.split('reader/')[0].concat(vol);
    link=link.concat(cap);
    if(page<10){
        link=link.concat("0");
    }
    link=link.concat(page+'');
    link=link.concat(".jpg");

    return link
  }

  private async pageExist(html: string, manga: string, chapter: string, page: number) {
    return (
      (
        await axios.get(
          this.pageUrl(html, manga, chapter, page),
          {validateStatus: () => true}
        )
      ).status === 200
    )
  }

}

export default new Juinjutsu()