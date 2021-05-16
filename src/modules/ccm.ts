import { SearchResult, Manga, Chapter } from '../classes/interfaces'
import Module from '../classes/Module'
import axios from 'axios'

interface ccmSearch {
  id: string
  title: string
}

interface ccmManga {
  chapters: {
    title: string
    volume: string
    chapter: string
    time: string
  }[]
  author: string
  artist: string
  cover: string
}

class CCM implements Module {
  id = 'ccm'
  name = 'CCM Translations'

  search(query: string) {
    return new Promise((resolve: (value: SearchResult[]) => void) => {
      axios.get(`https://api.ccmscans.in/mangas?title=${query}`).then((response) => {
        let result: SearchResult[] = [];

        response.data.forEach((el: ccmSearch) => {
          result.push({ id: el.id, title: el.title})
        })
        resolve(result)
      })
    })
  }

  manga(mangaid: string) {
    return new Promise((resolve: (value: Manga) => void) => {
      axios.get(`https://api.ccmscans.in/manga/${mangaid}`).then((response) => {
        let el: ccmManga = response.data

        let chapters: Chapter[] = []
        el.chapters.reverse().forEach((ch) => {
          chapters.push({ title: (ch.volume!='' ? 'Vol.'+ch.volume+' ' : '')+'Ch.'+ch.chapter+(ch.title!='' ? ' - '+ch.title : ''), id: ch.chapter, date: new Date(ch.time) })
        });

        resolve({ synopsis: '', author: el.author, artist: el.artist, img: el.cover, chapters: chapters, sourceurl: `https://ccmscans.in/manga/${mangaid}`})
      })
    })
  }

  chapter(manga: string, id: string) {
    return new Promise((resolve: (value: string[]) => void) => {
      axios.get(`https://api.ccmscans.in/chapter/${manga}/${id}`).then((response) => {
        resolve(response.data.images)
      })
    })
  }

  /*latest() {
    const promise = new Promise((resolve: (value: SearchResult[]) => void) => {
      axios.get(`https://api.ccmscans.in/chapters?sort=time&limit=20&grouped=1`).then((response) => {
        let result: SearchResult[] = [];

        response.data.forEach((el: ccmSearch) => {
          result.push({ id: el.id, title: el.title})
        })
        resolve(result)
      })
    })
    return promise
  }*/
}

export default new CCM()