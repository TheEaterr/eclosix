<script lang="ts">
	import { IconArrowLeft } from '@tabler/icons-svelte';

	let button: HTMLButtonElement | null = $state(null);
	const animate = () => {
		button?.animate(
			[
				{ transform: 'translate(0px)' },
				{ transform: 'translate(-5px)' },
				{ transform: 'translate(0px)' }
			],
			{
				duration: 500,
				iterations: Infinity,
				easing: 'ease-in-out'
			}
		);
	};
</script>

<div class="tooltip tooltip-bottom" data-tip="Retour">
	<button
		class="link"
		bind:this={button}
		onclick={() => {
			animate();
			window.history.back();
		}}
	>
		<IconArrowLeft class="swap-on h-10 w-10" />
	</button>
</div>
