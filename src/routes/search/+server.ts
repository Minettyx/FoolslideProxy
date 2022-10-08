import { ModuleFlags } from '$lib/server/classes/Module';
import { modules } from '$lib/server/modules';
import { compute_searchresult } from '$lib/server/utils';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request }) => {
	const body = await request.formData();

	let response = '';
	const search = (body.get('search') + '').toLowerCase().trim();

	// check if the search is general or specific
	const specific = (() => {
		for (const mod of modules) {
			if (search.startsWith(mod.id.toLowerCase() + ':')) {
				return mod;
			}
		}
		return false;
	})();

	if (specific) {
		const query = search.split(specific.id.toLowerCase())[1].substring(1).trim();
		const data = await specific.search(query);

		for (const ele of data) {
			response += `<div class="group"><div class="title"><a href="/series/${
				compute_searchresult(ele, specific.id)[0].manga_uid
			}" title="${ele.title}">${ele.title}</a></div></div>`;
		}
	} else {
		for (const mod of modules) {
			if (mod.flags.includes(ModuleFlags.DISABLE_GLOBAL_SEARCH)) continue;

			const data = await mod.search(search);

			for (const ele of data) {
				response += `<div class="group"><div class="title"><a href="/series/${
					compute_searchresult(ele, mod.id)[0].manga_uid
				}" title="${ele.title}">${ele.title}</a></div></div>`;
			}
		}
	}

	return new Response(response, {
		headers: { 'content-type': 'text/html', "Cache-Control": "max-age=3600, public" }
	});
};
