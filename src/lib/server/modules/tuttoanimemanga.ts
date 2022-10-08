import type { SearchResult, Manga } from '../classes/interfaces';
import Module from '../classes/Module';
import axios from 'axios';

class TuttoAnimeManga extends Module {
	id = 'tam';
	name = 'TuttoAnimeManga';

	search(query: string): Promise<SearchResult[]> {
		return new Promise(async (resolve: (value: SearchResult[]) => void) => {
			if (query.length < 3 && query !== '') {
				resolve([]);
			}

			const page = await axios.get(
				query === ''
					? 'https://tuttoanimemanga.net/api/comics'
					: 'https://tuttoanimemanga.net/api/search/' + encodeURIComponent(query)
			);

			resolve(
				page.data.comics.map((v: { url: string; title: string }) => ({
					id: v.url,
					title: v.title
				}))
			);
		});
	}

	manga(id: string): Promise<Manga> {
		return new Promise(async (resolve: (value: Manga) => void) => {
			const page = await axios.get('https://tuttoanimemanga.net/api' + id);
			const comic = page.data.comic;

			resolve({
				synopsis: comic.description,
				author: comic.author,
				artist: comic.artist,
				img: comic.thumbnail,
				chapters: comic.chapters.map(
					(v: { full_title: string; url: string; published_on: string }) => ({
						title: v.full_title,
						id: v.url,
						date: new Date(v.published_on)
					})
				),
				sourceurl: 'http://tuttoanimemanga.net' + id
			});
		});
	}
	chapter(manga: string, id: string): Promise<string[]> {
		return new Promise(async (resolve: (value: string[]) => void) => {
			const page = await axios.get('https://tuttoanimemanga.net/api' + id);

			resolve(page.data.chapter.pages);
		});
	}
}

export default new TuttoAnimeManga();
