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

			if (!buffer) {
				return new Response('404', {
					status: 404
				});
			}

			return new Response(buffer, {
				headers: { 'content-type': 'image/jpeg', 'Cache-Control': 'max-age=3600, public' }
			});
		}
	}

	return new Response('404', {
		status: 404
	});
};
