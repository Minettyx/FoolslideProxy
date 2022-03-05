import { SearchResult, Manga, Chapter } from '../classes/interfaces'
import Module from '../classes/Module'
import axios from 'axios'
import cheerio from 'cheerio'
import { NodeVM } from 'vm2'

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

      const baseUrl = await this.pageBaseUrl(html, manga, id)

      /** calculate oversize */

      let oversize = 15
      while (true) {
        if(await this.pageExist(baseUrl, oversize+1)) {
          oversize *= 2
        } else {
          break
        }
      }

      /** bynary search */

      let start = 1;
      let end = oversize;
      let size = oversize

      while (start <= end) {
        let middle = Math.floor((start + end) / 2);
  
        if (await this.pageExist(baseUrl, middle)) {
          if(end-start <= 1) {
            size = middle
            break
          }
          start = middle + 1
        } else {
          if(end-start <= 1) {
            size = middle-1
            break
          }
          end = middle - 1
        }
      }

      /** return results */

      let res: string[] = []
      for (let i = 1; i <= size; i++) {
        res.push(this.pageUrl(baseUrl, i))
      }

      resolve(res)

    })
  }

  private pageBaseUrl(html: string, manga: string, chapter: string) {
    return new Promise<string>((resolve) => {
      let url = new URL(chapter, 'http://onepiecepower.info/manga/'+manga).href
  
      let code =
        (
          html.split('<script type="text/javascript">').pop() || ''
        )
        .split('link=link.concat(".jpg")')[0]
        .split('\n')
        .filter(v => 
          !(
            v.trim().startsWith('$') ||
            v.trim().startsWith('if') ||
            v.trim().startsWith('}') ||
            v.includes('XMLHttpRequest')
          )
        )
        .join('\n')
        .split('window.location.href')
        .join("'"+addslashes(url)+"'")
        .split('location.pathname')
        .join("'"+addslashes(new URL(url).pathname)+"'")
        .split('location.search')
        .join("''")
  
      code = code.concat('\nmodule.exports = link;')

      const s = new NodeVM()
      const res = new URL('./', (s.run(code)+'')).href
      // console.log(res)
      resolve(res)
    })
  }

  private pageUrl(baseurl: string, page: number) {
    let link = baseurl
    if(page<10){
      link=link.concat("0");
    }
    link=link.concat(page+'');
    link=link.concat(".jpg");

    return link
  }

  private async pageExist(baseurl: string, page: number) {
    return (
      (
        await axios.get(
          this.pageUrl(baseurl, page),
          {validateStatus: () => true}
        )
      ).status === 200
    )
  }

}

export default new Juinjutsu()

function addslashes( str: string ) {
  return (str + '').replace(/[\\"']/g, '\\$&').replace(/\u0000/g, '\\0');
}