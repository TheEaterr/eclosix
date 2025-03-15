<script lang="ts">
	import Header from '$lib/components/header/Header.svelte';
</script>

<svelte:head>
	<title>À propos - Éclosix</title>
	<meta
		name="description"
		content="Comment Éclosix, un jeu quotidien où il faut écrire 12 mots avec 7 lettres, a été créé."
	/>
</svelte:head>

<Header reset={undefined} showHowToPlay={false} />

<div class="w-full text-center">
	<article class="prose bg-base-100 m-5 inline-block max-w-[120ch]">
		<h3 class="small-title text-primary text-3xl font-bold">À propos d'Éclosix</h3>
		<hr class="bg-neutral h-[1px] rounded" />
		<div class="text-justify">
			<p>TODO.</p>
		</div>
	</article>
</div>
