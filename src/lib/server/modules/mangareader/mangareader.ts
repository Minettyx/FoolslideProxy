import type { SearchResult, Manga, Chapter } from '../../classes/interfaces';
import Module, { ModuleFlags } from '../../classes/Module';
import axios from 'axios';
import { load as cheerioload } from 'cheerio';
import { ImageUnscrambler } from './ImageInterceptor';

class MangaReader extends Module {
	id = 'mr';
	name = 'MangaReader';
	flags: ModuleFlags[] = [ModuleFlags.DISABLE_GLOBAL_SEARCH];

	search(query: string) {
		return new Promise(async (resolve: (value: SearchResult[]) => void) => {
			if (query.length < 1) {
				resolve([]);
				return;
			}

			const $ = cheerioload(
				(await axios.get(`https://mangareader.to/search?keyword=${query}`)).data
			);

			const results: SearchResult[] = [];
			$('#main-content > section > div.manga_list-sbs > div.mls-wrap')
				.find('div.item.item-spc')
				.each((_, el) => {
					const ele = $(el);
					const a = ele.find('div.manga-detail > h3 > a');
					const langs = ele.find('a.manga-poster > span').text();

					langs.split('/').forEach((l) => {
						results.push({ id: a.attr('href') + '|' + l, title: `[${l}] ${a.attr('title')}` });
					});
				});

			resolve(results);
		});
	}

	manga(mangaid: string) {
		return new Promise(async (resolve: (value: Manga) => void) => {
			const { mangaId, lang } = (() => {
				const p = mangaid.split('|');
				return {
					mangaId: p[0],
					lang: p[1].toLowerCase()
				};
			})();

			const $ = cheerioload((await axios.get(`https://mangareader.to${mangaId}`)).data);

			const chapters: Chapter[] = [];
			$('#' + lang + '-chapters')
				.find('a')
				.each((_, ch) => {
					const chap = $(ch);
					chapters.push({
						title: chap.attr('title') + '',
						id: chap.attr('href') + '',
						date: new Date(0)
					});
				});

			resolve({
				synopsis: $('#modaldesc > div > div > div.modal-body > div').text().trim(),
				author:
					$(
						'#ani_detail > div > div > div.anis-content > div.anisc-detail > div.sort-desc > div.anisc-info-wrap > div.anisc-info > div:nth-child(3)'
					)
						.text()
						.split('Authors:')[1] || 'unknown',
				artist: '',
				img:
					$('#ani_detail > div > div > div.anis-content > div.anisc-poster > div > img').attr(
						'src'
					) + '',
				chapters,
				sourceurl: `https://mangareader.to${mangaId}`
			});
		});
	}

	chapter(manga: string, id: string) {
		return new Promise(async (resolve: (value: string[]) => void) => {
			let $ = cheerioload((await axios.get(`https://mangareader.to${id}`)).data);
			const chapid = $('#wrapper').attr('data-reading-id') + '';

			$ = cheerioload(
				(
					await axios.get(
						`https://mangareader.to/ajax/image/list/chap/${chapid}?mode=horizontal&quality=high&hozPageSize=1`
					)
				).data.html + ''
			);

			const res: string[] = [];
			const images = $('#divslide > div.divslide-wrapper').find('.ds-image');
			images.each((i, el) => {
				if (i == images.length - 1) return;
				const ele = $(el);
				if (ele.hasClass('shuffled')) {
					res.push(this.imageProxyUrl('', '', ele.attr('data-url') + ''));
				} else {
					res.push(ele.attr('data-url') + '');
				}
			});

			resolve(res);

			// console.log(res)
		});
	}

	image(manga: string, chapter: string, id: string) {
		return new Promise(async (resolve: (value: Buffer) => void) => {
			// console.log(manga, chapter, id)

			const response = await axios.get(id, {
				responseType: 'arraybuffer'
			});

			const init = new Uint8Array(Buffer.from(response.data, 'binary'));

			const res = await new ImageUnscrambler().interceptResponse(init);

			resolve(res);
		});
	}
}

export default new MangaReader();
