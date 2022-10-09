<script lang="ts">
	import type { PageData } from './$types';
	import 'github-markdown-css/github-markdown-light.css';

	export let data: PageData;
</script>

<article class="markdown-body">
	{@html data.html}
</article>

<style>
	.markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
	}

	@media (max-width: 767px) {
		.markdown-body {
			padding: 15px;
		}
	}
</style>
