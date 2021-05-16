interface SearchResult {
  id: string
  title: string
}

interface Manga {
  synopsis: string
  author: string
  artist: string
  img: string
  chapters: Chapter[]
  sourceurl: string
}

interface Chapter {
  title: string
  id: string
  date: Date
}

export { SearchResult, Manga, Chapter }