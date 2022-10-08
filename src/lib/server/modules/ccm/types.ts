export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
	ID: string;
	String: string;
	Boolean: boolean;
	Int: number;
	Float: number;
};

export type Manga = {
	__typename?: 'Manga';
	id: Scalars['String'];
	title: Scalars['String'];
	author: Scalars['String'];
	artist?: Maybe<Scalars['String']>;
	status: Scalars['Int'];
	cover: Scalars['String'];
	chapters: Array<Chapter>;
};

export type Chapter = {
	__typename?: 'Chapter';
	chapter: Scalars['String'];
	manga: Manga;
	title: Scalars['String'];
	volume: Scalars['String'];
	time: Scalars['String'];
	webtoon: Scalars['Boolean'];
	images: Array<Scalars['String']>;
};

export type Query = {
	__typename?: 'Query';
	chapter?: Maybe<Chapter>;
	chapters: Array<Chapter>;
	manga?: Maybe<Manga>;
	mangas: Array<Manga>;
};

export type QueryChapterArgs = {
	manga: Scalars['String'];
	chapter: Scalars['String'];
};

export type QueryChaptersArgs = {
	limit?: Maybe<Scalars['Int']>;
};

export type QueryMangaArgs = {
	id: Scalars['String'];
};

export type QueryMangasArgs = {
	search?: Maybe<Scalars['String']>;
};
