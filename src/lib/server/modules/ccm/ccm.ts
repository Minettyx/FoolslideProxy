import type { SearchResult, Manga } from '../../classes/interfaces';
import Module from '../../classes/Module';
import axios from 'axios';
import type { Manga as CCMManga, Chapter as CCMChapter } from './types';

const BASEURL = 'https://ccmscans.in/api/';

class CCM extends Module {
	id = 'ccm';
	name = 'CCM Translations';

	search(query: string) {
		return new Promise(async (resolve: (value: SearchResult[]) => void) => {
			const data: CCMManga[] = (await axios.get(BASEURL + 'mangas.json')).data;

			const result: SearchResult[] = [];
			for (const manga of data) {
				if (manga.title.toLocaleLowerCase().includes(query.toLocaleLowerCase())) {
					result.push({
						id: manga.id,
						title: manga.title
					});
				}
			}
			resolve(result);
		});
	}

	manga(mangaid: string) {
		return new Promise(async (resolve: (value: Manga) => void) => {
			const data: CCMManga = (await axios.get(BASEURL + `manga/${mangaid}.json`)).data;

			resolve({
				synopsis: '',
				author: data.author,
				artist: data.artist || '',
				img: data.cover,
				chapters: data.chapters.reverse().map((c) => {
					return {
						title:
							(c.volume != '' ? 'Vol.' + c.volume + ' ' : '') +
							'Ch.' +
							c.chapter +
							(c.title != '' ? ' - ' + c.title : ''),
						id: c.chapter,
						date: new Date(c.time)
					};
				}),
				sourceurl: `https://ccmscans.in/manga/${data.id}`
			});
		});
	}

	chapter(manga: string, id: string) {
		return new Promise(async (resolve: (value: string[]) => void) => {
			const data: CCMChapter = (await axios.get(BASEURL + `chapter/${manga}/${id}.json`)).data;

			resolve(data.images);
		});
	}
}

export default new CCM();
