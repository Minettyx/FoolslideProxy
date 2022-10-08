import { modules } from '$lib/server/modules';
import { compute_chapter } from '$lib/server/utils';
import { authorartist, fromhex } from '$lib/utils';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ params }) => {
	const modid = params.mod;
	const mangaid = fromhex(params.manga);

	for (const mod of modules) {
		if (mod.id === modid) {
			const data = await mod.manga(mangaid);

			let response = `<html><head></head><body><div id="wrapper"><article id="content"><div class="panel"><div class="comic info"><div class="thumbnail"><img src="${
				data.img
			}" /></div><div class="large comic"><h1 class="title"></h1><div class="info"><b>Author</b>: ${authorartist(
				data.author,
				data.artist
			)}<br><b>Artist</b>: ${mod.name}<br><b>Synopsis</b>: ${
				data.synopsis
			}</div></div></div><div class="list"><div class="group"><div class="title">Volume</div>`;

			data.chapters.forEach((chapter) => {
				response += `<div class="element"><div class="title"><a href="/read/${
					compute_chapter(chapter, mangaid, mod.id)[0].chapter_uid
				}" title="${chapter.title}">${
					chapter.title
				}</a></div><div class="meta_r">by <a href="" title="" ></a>, ${chapter.date.getFullYear()}.${(
					'0' +
					(chapter.date.getMonth() + 1)
				).slice(-2)}.${chapter.date.getDate()}</div></div>`;
			});

			response += `</div></div></div></article></div></body></html>`;

			return new Response(response, {
				headers: { 'content-type': 'text/html' }
			});
		}
	}

	return new Response('404', {
		status: 404
	});
};

export const GET: RequestHandler = async ({ params }) => {
	const modid = params.mod;
	const mangaid = fromhex(params.manga);

	for (const mod of modules) {
		if (mod.id === modid) {
			const data = await mod.manga(mangaid);

			return new Response('', {
				headers: { Location: data.sourceurl },
				status: 302
			});
		}
	}

	return new Response('404', {
		status: 404
	});
};
