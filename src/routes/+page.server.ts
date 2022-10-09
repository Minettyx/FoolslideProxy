import pkg from 'showdown';
const { Converter } = pkg;
import type { PageServerLoad } from './$types';
import { promises } from 'fs';

export const load: PageServerLoad = async () => {
	const converter = new Converter();

	const md = await promises.readFile('README.md', 'utf8');
	const html = converter.makeHtml(md);

	return { html };
};
