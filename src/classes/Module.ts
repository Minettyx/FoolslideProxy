import { SearchResult, Manga } from './interfaces'

interface Module {
  id: string
  name: string
  
  search(query: string):Promise<SearchResult[]>
  manga(id: string):Promise<Manga>
  chapter(manga: string, id: string):Promise<string[]>
}

export default Module