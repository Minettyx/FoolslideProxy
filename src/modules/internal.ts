import { SearchResult, Manga, Chapter } from '../classes/interfaces'
import Module from '../classes/Module'

class Internal implements Module {
  id: string = 'internal'
  name: string = 'Internal'

  search(_: string): Promise<SearchResult[]> {
    return new Promise(resolve => {
      resolve([])
    })
  }

  manga(id: string): Promise<Manga> {
    return new Promise(resolve => {
      if(id === 'supportedsources') {

        let syn = 'Click on WebView above';

        const manga: Manga = {
          synopsis: syn,
          author: 'Minettyx',
          artist: '',
          img: '',
          chapters: [{
            title: 'Click on WebView above',
            id: '',
            date: new Date('27/02/2022')
          }],
          sourceurl: 'https://github.com/Minettyx/FoolslideProxy/tree/master/docs/sources.md',
        }

        resolve(manga)

      }
    })
  }

  chapter(): Promise<string[]> {
    return new Promise(resolve => {
      resolve([''])
    })
  }
    
}

export default new Internal()