import { SearchResult, Manga, Chapter } from '../../classes/interfaces'
import Module from '../../classes/Module'
import { createClient, OperationResult } from '@urql/core'
import { Query } from './types'

import 'isomorphic-unfetch'
const client = createClient({
  url: 'https://api.ccmscans.in/graphql',
});

class CCM extends Module {
  id = 'ccm'
  name = 'CCM Translations'

  search(query: string) {
    return new Promise((resolve: (value: SearchResult[]) => void) => {
      client.query(`
        query($title: String!) {
          mangas(search: $title) {
            title
            id
          }
        }
      `, { title: query }).toPromise().then((response: OperationResult<Query, { title: string; }>) => {
        // console.log(response.data?.mangas)
        let result: SearchResult[] = [];

        response.data?.mangas.forEach((el) => {
          result.push({ id: el.id, title: el.title })
        })
        resolve(result)
      })
    })
  }

  manga(mangaid: string) {
    return new Promise((resolve: (value: Manga) => void) => {
      client.query(`
        query ($id: String!) {
          manga(id: $id) {
            id
            cover
            author
            artist
            chapters {
              chapter
              volume
              title
              time
            }
          }
        }
      `, { id: mangaid }).toPromise().then((response: OperationResult<Query, { id: string; }>) => {
        let el = response.data?.manga
        if (!el) return

        let chapters: Chapter[] = []
        el.chapters.reverse().forEach((ch) => {
          chapters.push({ title: (ch.volume != '' ? 'Vol.' + ch.volume + ' ' : '') + 'Ch.' + ch.chapter + (ch.title != '' ? ' - ' + ch.title : ''), id: ch.chapter, date: new Date(parseInt(ch.time)) })
        });

        resolve({ synopsis: '', author: el.author, artist: el.artist || '', img: el.cover, chapters: chapters, sourceurl: `https://ccmscans.in/manga/${el.id}` })
      })
    })
  }

  chapter(manga: string, id: string) {
    return new Promise((resolve: (value: string[]) => void) => {
      client.query(`
        query($manga: String!, $chapter: String!) {
          chapter(manga: $manga, chapter: $chapter) {
            images
          }
        }
      `, {manga, chapter: id}).toPromise().then((response: OperationResult<Query, { manga: string; chapter: string; }>) => {
        resolve(response.data?.chapter?.images||[])
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