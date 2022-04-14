import { SearchResult, Manga, Chapter } from '../classes/interfaces'
import Module, { ModuleFlags } from '../classes/Module'
import { modules } from '../main'

class Internal extends Module {
  id: string = 'internal'
  name: string = 'Internal'
  flags: ModuleFlags[] = [
    ModuleFlags.HIDDEN,
    ModuleFlags.DISABLE_GLOBAL_SEARCH
  ]

  search(_: string): Promise<SearchResult[]> {
    return new Promise(resolve => {
      resolve([])
    })
  }

  manga(id: string): Promise<Manga> {
    return new Promise(resolve => {
      if(id === 'supportedsources') {

        let syn = 'Click on WebView for more infos';

        let list = []
        for(const mod of modules.slice().reverse()) {
          if(mod.flags.includes(ModuleFlags.HIDDEN)) continue

          let title = `${mod.name} (${mod.id})`

          list.push({
            title,
            id: mod.id,
            date: new Date(0)
          })
        }

        const manga: Manga = {
          synopsis: syn,
          author: 'Minettyx',
          artist: '',
          img: '',
          chapters: list,
          sourceurl: 'https://github.com/Minettyx/FoolslideProxy/tree/master/docs/sources.md',
        }

        resolve(manga)

      }
    })
  }

  chapter(): Promise<string[]> {
    return new Promise(resolve => {
      resolve([])
    })
  }
    
}

export default new Internal()