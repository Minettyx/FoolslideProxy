import express from 'express'
import apicache, { id } from 'apicache'
import Module from './classes/Module'
const router = express.Router()

import ccm from './modules/ccm'
import mangaworld from './modules/mangaworld'
import juinjutsu from './modules/juinjutsu'

const cache = apicache.middleware;



/** Initialize Modules */
const modules: Module[] = [
  ccm,
  mangaworld,
  juinjutsu
]





router.post('/search', (req, res, next) => {
  let response = "";
  let c = 0

  modules.forEach(async (mod) => {

    const search = req.body.search ? req.body.search as string : ''

    if(!search.includes(':') || search.toLowerCase().startsWith(mod.id)) {
      const query = search.includes(':') ? search.split(mod.id)[1].substring(1).trim() : search
      const data = await mod.search(query)

      data.forEach(ele => {
        response += `<div class="group"><div class="title"><a href="/series/${mod.id}-${btoa(ele.id)}" title="${ele.title}">${ele.title}</a></div></div>`
      })

      c++ ;if(c === modules.length) res.send(response)
    } else {
      c++ ;if(c === modules.length) res.send(response)
    }
    
  })

})

router.post('/series/:id', (req, res, next) => {

  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')

  modules.forEach(async (mod) => {
    if(mod.id === modid) {
      mod.manga(mangaid).then((data) => {

        let response = `<html><head></head><body><div id="wrapper"><article id="content"><div class="panel"><div class="comic info"><div class="thumbnail"><img src="${data.img}" /></div><div class="large comic"><h1 class="title"></h1><div class="info"><b>Author</b>: ${authorartist(data.author, data.artist)}<br><b>Artist</b>: ${mod.name}<br><b>Synopsis</b>: ${data.synopsis}</div></div></div><div class="list"><div class="group"><div class="title">Volume</div>`

        data.chapters.forEach(chapter => {
          response += `<div class="element"><div class="title"><a href="/read/${mod.id}-${btoa(mangaid)}-${btoa(chapter.id)}" title="${chapter.title}">${chapter.title}</a></div><div class="meta_r">by <a href="" title="" ></a>, ${chapter.date.getFullYear()}.${("0"+(chapter.date.getMonth()+1)).slice(-2)}.${chapter.date.getDate()}</div></div>`
        })
          
        response += `</div></div></div></article></div></body></html>`
        res.send(response)
      })
    }
  })
})

router.get('/series/:id', (req, res, next) => {
  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')

  modules.forEach(async (mod) => {
    if(mod.id === modid) {
      mod.manga(mangaid).then((data) => {
        res.redirect(data.sourceurl)
      })
    }
  })
})

router.post('/read/:id', (req, res, next) => {

  const modid = req.params.id.split("-")[0]
  const mangaid = atob(req.params.id.split("-")[1]+'')
  const id = atob(req.params.id.split("-")[2]+'')

  modules.forEach(async (mod) => {
    if(mod.id === modid) {
      mod.chapter(mangaid, id).then((images) => {
        let data: {url: string}[] = []
        images.forEach(img => {
          data.push({url: img});
        });

        res.send("var pages = "+JSON.stringify(data)+";");
      })
    }
  })

})

router.get('/directory/1/', (req, res, next) => {
  res.send(`<div class="group"><div class="title"><a href="/series/" title="">Latest feed not supported</a></div></div>`)
})


export = router


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