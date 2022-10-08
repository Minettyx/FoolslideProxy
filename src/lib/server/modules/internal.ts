import type { SearchResult, Manga } from '../classes/interfaces';
import Module, { ModuleFlags } from '../classes/Module';
import { modules } from '../modules';

class Internal extends Module {
	id = 'internal';
	name = 'Internal';
	flags: ModuleFlags[] = [ModuleFlags.HIDDEN, ModuleFlags.DISABLE_GLOBAL_SEARCH];

	search(): Promise<SearchResult[]> {
		return new Promise((resolve) => {
			resolve([]);
		});
	}

	manga(id: string): Promise<Manga> {
		return new Promise((resolve) => {
			if (id === 'supportedsources') {
				const syn = 'Click on WebView for more infos';

				const list = [];
				for (const mod of modules.slice().reverse()) {
					if (mod.flags.includes(ModuleFlags.HIDDEN)) continue;

					const title = `${mod.name} (${mod.id})`;

					list.push({
						title,
						id: mod.id,
						date: new Date(0)
					});
				}

				const manga: Manga = {
					synopsis: syn,
					author: 'Minettyx',
					artist: '',
					img: '',
					chapters: list,
					sourceurl: 'https://github.com/Minettyx/FoolslideProxy/wiki/Available-sources'
				};

				resolve(manga);
			}
		});
	}

	chapter(): Promise<string[]> {
		return new Promise((resolve) => {
			resolve([]);
		});
	}
}

export default new Internal();
