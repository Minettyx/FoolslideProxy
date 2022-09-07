import express from 'express'
import Module, { ModuleFlags } from './classes/Module'
const router = express.Router()

import internal from './modules/internal'

import ccm from './modules/ccm/ccm'
import mangaworld from './modules/mangaworld'
import juinjutsu from './modules/juinjutsu'
import onepiecepower from './modules/onepiecepower'
import tuttoanimemanga from './modules/tuttoanimemanga'
import mangareader from './modules/mangareader/mangareader'



/** Initialize Modules */
export const modules: ReadonlyArray<Module> = [
  internal,
  ccm,
  mangaworld,
  juinjutsu,
  onepiecepower,
  tuttoanimemanga,
  mangareader
]




router.post('/search', async (req, res, next) => {
  let response = ''
  const search = (req.body.search+'').toLowerCase().trim()
  
  // check if the search is general or specific
  const specific = (() => {
    for (const mod of modules) {
      if(search.startsWith(mod.id.toLowerCase()+':')) {
        return mod
      }
    }
    return false
  })()

  if(specific) {
    const query = search.split(specific.id.toLowerCase())[1].substring(1).trim()
    const data = await specific.search(query)

    for(const ele of data) {
      response += `<div class="group"><div class="title"><a href="/series/${specific.id}-${btoa(ele.id)}" title="${ele.title}">${ele.title}</a></div></div>`
    }
  } else {
    for (const mod of modules) {
      if(mod.flags.includes(ModuleFlags.DISABLE_GLOBAL_SEARCH)) continue

      const data = await mod.search(search)

      for(const ele of data) {
        response += `<div class="group"><div class="title"><a href="/series/${mod.id}-${btoa(ele.id)}" title="${ele.title}">${ele.title}</a></div></div>`
      }
    }
  }

  res.send(response)
  return

})

router.post('/series/:id', async (req, res, next) => {

  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')

  for(const mod of modules) {
    if(mod.id === modid) {
      const data = await mod.manga(mangaid)

      let response = `<html><head></head><body><div id="wrapper"><article id="content"><div class="panel"><div class="comic info"><div class="thumbnail"><img src="${data.img}" /></div><div class="large comic"><h1 class="title"></h1><div class="info"><b>Author</b>: ${authorartist(data.author, data.artist)}<br><b>Artist</b>: ${mod.name}<br><b>Synopsis</b>: ${data.synopsis}</div></div></div><div class="list"><div class="group"><div class="title">Volume</div>`

      data.chapters.forEach(chapter => {
        response += `<div class="element"><div class="title"><a href="/read/${mod.id}-${btoa(mangaid)}-${btoa(chapter.id)}" title="${chapter.title}">${chapter.title}</a></div><div class="meta_r">by <a href="" title="" ></a>, ${chapter.date.getFullYear()}.${("0"+(chapter.date.getMonth()+1)).slice(-2)}.${chapter.date.getDate()}</div></div>`
      })
        
      response += `</div></div></div></article></div></body></html>`
      res.send(response)
      return
    }
  }
})

router.get('/series/:id', async (req, res, next) => {
  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')

  for(const mod of modules) {
    if(mod.id === modid) {
      const data = await mod.manga(mangaid)
      res.redirect(data.sourceurl)
      return
    }
  }
})

router.post('/read/:id', async (req, res, next) => {

  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')
  const id = atob(req.params.id.split("-")[2]+'')

  for(const mod of modules) {
    if(mod.id === modid) {
      const images = await mod.chapter(mangaid, id)

      const data: {url: string}[] = images.map(v => ({url: v}))

      res.send("<script>var pages = "+JSON.stringify(data)+";</script>");
      return
    }
  }

})

router.get('/directory/1/', (_, res, next) => {
  res.send(`<div class="group"><div class="title"><a href="/series/internal-${btoa('supportedsources')}" title="">Supported sources</a></div></div>`)
})

router.get('/image/:id', async (req, res, next) => {
  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')
  const chapterid = atob(req.params.id.split("-")[2]+'')
  const imageid = atob(req.params.id.split("-")[3]+'')

  for(const mod of modules) {
    if(mod.id === modid && mod.image) {
      const buffer = await mod.image(mangaid, chapterid, imageid)

      res.contentType('image/jpeg');
      res.send(buffer)
      res.end()
      return
    }
  }
})


export default router


function atob(hex: string) {
  return Buffer.from(hex, 'hex').toString()
}

function btoa(string: string) {
  return Buffer.from(string).toString('hex')
}

function authorartist(author: string, artist: string): string {
  let names = []
  if(author !== '') names.push(author)
  if(artist !== '') names.push(artist)

  if(names.length > 1) {
    if(names[0] === names[1]) {
      return names[0]
    } else {
      return names[0]+', '+names[1]
    }
  } else if(names.length === 1) {
    return names[0]
  }

  return ''
}
