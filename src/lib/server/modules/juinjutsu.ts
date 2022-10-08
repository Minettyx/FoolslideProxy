import type { SearchResult, Manga, Chapter } from '../classes/interfaces';
import Module from '../classes/Module';
import axios from 'axios';
import { load as cheerioload } from 'cheerio';

class Juinjutsu extends Module {
	id = 'jj';
	name = 'JuinJutsu';

	search(query: string) {
		return new Promise(async (resolve: (value: SearchResult[]) => void) => {
			if (query.length < 2) {
				resolve([]);
				return;
			}

			const page1 = await axios.get('https://juinjutsureader.ovh/directory/1/');
			const parsed = cheerioload(page1.data);

			const pages_count = parseInt(
				parsed('.gbuttonright.btn.btn-primary').attr('href')?.split('directory/')[1].split('/')[0] +
					''
			);

			const pages = [];
			pages.push(page1.data);
			for (let i = 2; i <= pages_count; i++) {
				const p = await axios.get('https://juinjutsureader.ovh/directory/' + i + '/');
				pages.push(p.data);
			}

			const chs: { title: string; id: string }[] = [];
			pages.forEach((el) => {
				const $ = cheerioload(el);
				const caps = $('.series_element');
				caps.each((i, capEl) => {
					const cap = { title: '', id: '' };
					cap.title = $('.title', capEl).find('a').text();
					cap.id = $('.title', capEl).find('a').attr('href')?.split('series/')[1] + '';
					chs.push(cap);
				});
			});

			const results: { title: string; id: string }[] = [];
			chs.forEach((ele) => {
				if (ele.title.toLowerCase().includes(query.toLowerCase())) {
					results.push(ele);
				}
			});

			resolve(results);
		});
	}

	manga(mangaid: string) {
		return new Promise(async (resolve: (value: Manga) => void) => {
			const page = await axios.get('https://juinjutsureader.ovh/series/' + mangaid);
			const $ = cheerioload(page.data);

			const chapters: Chapter[] = [];
			const data = {
				img: '',
				synopsis: '',
				author: '',
				artist: '',
				sourceurl: 'https://juinjutsureader.ovh/series/' + mangaid,
				chapters: chapters
			};
			data.img = $('img.thumb').attr('src') + '';
			data.synopsis = $('.trama') ? $('.trama').text().substring(7) : '';
			data.author = $('.autore') ? $('.autore').text().substring(8) : '';
			data.artist = $('.artista') ? $('.artista').text().substring(9) : '';

			$('.element').each((i, el) => {
				const a = $(el).find('a');
				const chtitle = a.text();
				const chid = a.attr('href')?.split(mangaid).pop() + '';
				const chdate = psDate($(el).find('.meta_r').text());

				data.chapters.push({ title: chtitle, id: chid, date: chdate });
			});

			resolve(data);
		});
	}

	chapter(manga: string, id: string) {
		return new Promise(async (resolve: (value: string[]) => void) => {
			const page = await axios('https://juinjutsureader.ovh/read/' + manga + id);
			const json = JSON.parse(page.data.split('var pages = ')[1].split(';')[0]);

			const results: string[] = [];
			json.forEach((el: { url: string }) => {
				results.push(el.url);
			});
			resolve(results);
		});
	}
}

export default new Juinjutsu();

function psDate(date: string): Date {
	switch (date) {
		case 'Oggi':
			return new Date();
		case 'Ieri': {
			const now = new Date();
			now.setDate(now.getDate() - 1);
			return now;
		}
		default: {
			const dt = new Date();
			const p = date.split('.');
			dt.setFullYear(parseInt(p[0]));
			dt.setMonth(parseInt(p[1]) - 1);
			dt.setDate(parseInt(p[2]));
			return dt;
		}
	}
}
