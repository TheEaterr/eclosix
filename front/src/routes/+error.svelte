<script lang="ts">
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import { IconExclamationCircle, IconExternalLink } from '@tabler/icons-svelte';
</script>

<svelte:head>
	<title>Éclosix - Erreur {page.status}</title>
	<meta
		name="description"
		content="Éclosix est un jeu de lettre où le but est d'écrire 12 mots en n'utilisant que 7 lettres. Plus le mot est long, plus il rapporte de points !"
	/>
</svelte:head>

<main class="fixed top-0 flex min-h-screen w-screen content-center items-center justify-center">
	<div class="hero h-fit w-fit">
		<div class="hero-overlay bg-base-100"></div>
		<div class="hero-content p-10 text-center sm:p-14 lg:p-20">
			<div class="max-w-md">
				<h1 class="title text-error text-5xl font-bold sm:text-6xl lg:text-7xl">
					{page.status}
				</h1>
				<div class="m-4 text-center">
					<span role="alert" class="alert alert-error">
						<IconExclamationCircle size={30} class="stroke-current" />
						<span>
							{#if page.status === 404}
								Cette page n'existe pas.
							{:else if page.status === 500}
								Un problème est survenu avec Éclosix. Si le problème persiste, signalez-le sur le
								<a
									href="https://github.com/TheEaterr/eclosix"
									target="_blank"
									rel="noopener"
									class="link">dépôt github<IconExternalLink class="inline" size={15} /></a
								>.
							{:else}
								{page.error?.message}
							{/if}
						</span>
					</span>
				</div>
				<div>
					{#if page.status !== 404}
						<button onclick={() => location.reload()} class="btn-neutral-special btn btn-sm"
							>Recharger la page</button
						>
					{/if}
					<button onclick={() => goto('/')} class="btn-neutral-special btn btn-sm ml-1"
						>Retour à l'accueil</button
					>
				</div>
			</div>
		</div>
	</div>
</main>

<style>
	.title {
		-webkit-text-stroke: 2px oklch(var(--nc));
	}
</style>
