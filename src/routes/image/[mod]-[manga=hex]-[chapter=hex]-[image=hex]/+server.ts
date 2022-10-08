import { modules } from '$lib/server/modules';
import { fromhex } from '$lib/utils';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ params }) => {
	const modid = params.mod;
	const mangaid = fromhex(params.manga);
	const chapterid = fromhex(params.chapter);
	const imageid = fromhex(params.image);

	for (const mod of modules) {
		if (mod.id === modid && mod.image) {
			const buffer = await mod.image(mangaid, chapterid, imageid);

			return new Response(buffer, {
				headers: { 'content-type': 'image/jpeg' }
			});
		}
	}

	return new Response('', {
		status: 404
	});
};
