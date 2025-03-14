<script lang="ts">
	import { goto } from '$app/navigation';
	import Header from '$lib/components/header/Header.svelte';
</script>

<svelte:head>
	<title>Éclosix</title>
	<meta
		name="description"
		content="Éclosix est un jeu de lettre où le but est d'écrire 12 mots en n'utilisant que 7 lettres. Plus le mot est long, plus il rapporte de points !"
	/>
</svelte:head>

<Header />

<main class="fixed top-0 flex min-h-screen w-screen content-center items-center justify-center">
	<div class="hero h-fit w-fit">
		<div class="hero-overlay bg-base-100"></div>
		<div class="hero-content p-10 text-center sm:p-14 lg:p-20">
			<div class="max-w-md">
				<h1 class="title text-primary text-5xl font-bold sm:text-6xl lg:text-7xl">Éclosix</h1>
				<div class="m-3 mt-5">
					<button
						onclick={() => goto('/game/daily')}
						class="btn-primary-special sm:btn-lg btn text-lg">Quotidienne</button
					>
					<button
						onclick={() => goto('/game/endless')}
						class="btn-secondary-special sm:btn-lg btn ml-5 text-lg">Infini</button
					>
				</div>
				<div>
					<button onclick={() => goto('/about')} class="btn-neutral-special btn btn-sm ml-1"
						>À propos</button
					>
				</div>
			</div>
		</div>
	</div>
</main>

<style>
	.title {
		-webkit-text-stroke: 2px var(--color-neutral-content);
	}
</style>
