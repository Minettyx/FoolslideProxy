import { tohex } from '$lib/utils';
import type {
	Chapter,
	ComputedChapter,
	ComputedManga,
	ComputedSearchResult,
	Manga,
	SearchResult
} from './classes/interfaces';

export function compute_searchresult(
	search: SearchResult[] | SearchResult,
	module_id: string
): ComputedSearchResult[] {
	const results: ComputedSearchResult[] = [];

	for (const res of search instanceof Array ? search : [search]) {
		results.push({
			manga_uid: `${module_id}-${tohex(res.id)}`,
			title: res.title
		});
	}

	return results;
}

export function compute_chapter(
	chapter: Chapter[] | Chapter,
	manga_id: string,
	module_id: string
): ComputedChapter[] {
	const results: ComputedChapter[] = [];

	for (const ch of chapter instanceof Array ? chapter : [chapter]) {
		results.push({
			title: ch.title,
			chapter_uid: `${module_id}-${tohex(manga_id)}-${tohex(ch.id)}`,
			date: ch.date
		});
	}

	return results;
}

export function compute_manga(manga: Manga, manga_id: string, module_id: string): ComputedManga {
	const result: ComputedManga = {
		...manga,
		chapters: compute_chapter(manga.chapters, manga_id, module_id)
	};
	return result;
}
