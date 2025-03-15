<script lang="ts">
	import Header from '$lib/components/header/Header.svelte';
	import { IconExternalLink } from '@tabler/icons-svelte';
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
			<p>
				Éclosix est une version française du jeu <a
					href="https://www.merriam-webster.com/games/blossom-word-game"
					target="_blank"
					rel="noopener"
					class="link">blossom<IconExternalLink class="inline" size={15} /></a
				>
				crée par Merriam-Webster. Ce projet n'a aucune association ni connexion avec le jeu
				original. Le site est entièrement open-source et disponible sur
				<a
					href="https://github.com/TheEaterr/taxonomicle"
					target="_blank"
					rel="noopener"
					class="link">ce dépôt<IconExternalLink class="inline" size={15} /></a
				>.
			</p>
			<h3>Création des données</h3>
			<p>
				La liste de mot d'Éclosix est issue du dictionnaire Français-GUTenberg et a été récupérée
				sur <a
					href="https://github.com/chrplr/openlexicon/blob/master/datasets-info/Liste-de-mots-francais-Gutenberg/README-liste-francais-Gutenberg.md"
					target="_blank"
					rel="noopener"
					class="link">ce lien<IconExternalLink class="inline" size={15} /></a
				>. 
			</p>
			<p>
				Pour générer les problèmes, la liste des mots français contenant exactement 7 lettres est
				parcourue. Pour chacun d'entre-eux, si cet ensemble de lettres n'a pas été encore considéré,
				alors tous les mots français de plus de 4 lettres qu'il est possible de créer avec ces
				lettres sont générés. Enfin, une vérification pour chacune des lettres est fait pour
				garantir que pour une lettre centrale donnée, il existe 12 mots différents contenant cette lettre.
			</p>
			<h3>Architecture du site</h3>
			<p>
				Le site est construit avec Svelte et SvelteKit pour le frontend, PocketBase pour le backend
				et TailwindCSS et DaisyUI pour le design.
			</p>
		</div>
	</article>
</div>
