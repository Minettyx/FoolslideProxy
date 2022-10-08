import type Module from './classes/Module';

import internal from './modules/internal';

import ccm from './modules/ccm/ccm';
import mangaworld from './modules/mangaworld';
import juinjutsu from './modules/juinjutsu';
import onepiecepower from './modules/onepiecepower';
import tuttoanimemanga from './modules/tuttoanimemanga';
import mangareader from './modules/mangareader/mangareader';

/** Initialize Modules */
export const modules: ReadonlyArray<Module> = [
	internal,
	ccm,
	mangaworld,
	juinjutsu,
	onepiecepower,
	tuttoanimemanga,
	mangareader
];
