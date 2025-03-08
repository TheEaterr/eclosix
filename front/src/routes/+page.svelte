<script lang="ts">
	import { goto } from '$app/navigation';
	import Header from '$lib/components/header/Header.svelte';
</script>

<svelte:head>
	<title>Taxonomicle</title>
	<meta
		name="description"
		content="Taxonomicle is game where you have to find a species in the tree of life and learn some taxonomy facts along the way. A new species is available each day but you can also play in endless mode and share the challenge with a friend!"
	/>
</svelte:head>

<Header onlyTheme={true} />

<main class="fixed top-0 flex min-h-screen w-screen content-center items-center justify-center">
	<div class="hero h-fit w-fit">
		<div class="hero-overlay bg-base-100"></div>
		<div class="hero-content p-10 text-center sm:p-14 lg:p-20">
			<div class="max-w-md">
				<h1 class="title text-primary text-5xl font-bold sm:text-6xl lg:text-7xl">Eclosix</h1>
				<div class="m-3 mt-5">
					<button onclick={() => goto('/game/daily')} class="btn-primary-special main-btn btn"
						>Daily</button
					>
					<button
						onclick={() => goto('/game/endless')}
						class="btn-secondary-special main-btn btn ml-5">Endless</button
					>
				</div>
				<div>
					<button onclick={() => goto('/game/tutorial')} class="btn-neutral-special btn btn-sm"
						>How to play</button
					>
					<button onclick={() => goto('/about')} class="btn-neutral-special btn btn-sm ml-1"
						>About</button
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
