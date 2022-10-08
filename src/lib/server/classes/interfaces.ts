export type SearchResult = {
	id: string;
	title: string;
};

export type Manga = {
	synopsis: string;
	author: string;
	artist: string;
	img: string;
	chapters: Chapter[];
	sourceurl: string;
};

export type Chapter = {
	title: string;
	id: string;
	date: Date;
};

export type ComputedSearchResult = {
	manga_uid: string;
	title: string;
};

export type ComputedManga = {
	synopsis: string;
	author: string;
	artist: string;
	img: string;
	chapters: ComputedChapter[];
	sourceurl: string;
};

export type ComputedChapter = {
	title: string;
	chapter_uid: string;
	date: Date;
};
