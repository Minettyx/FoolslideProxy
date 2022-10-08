import { tohex } from '$lib/utils';
import type { SearchResult, Manga } from './interfaces';

export default abstract class Module {
	abstract id: string;
	abstract name: string;
	flags: ModuleFlags[] = [];

	abstract search(query: string, language?: string): Promise<SearchResult[]>;
	abstract manga(id: string): Promise<Manga>;
	abstract chapter(manga: string, id: string): Promise<string[]>;
	public image?(manga: string, chapter: string, id: string): Promise<Buffer>;

	protected imageProxyUrl(manga: string, chapter: string, id: string): string {
		return (
			'/image/' +
			this.id +
			'-' +
			(manga === '' ? '0' : tohex(manga)) +
			'-' +
			(chapter === '' ? '0' : tohex(chapter)) +
			'-' +
			tohex(id)
		);
	}
}

export enum ModuleFlags {
	DISABLE_GLOBAL_SEARCH,
	HIDDEN
}
