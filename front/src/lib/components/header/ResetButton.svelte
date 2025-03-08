<script lang="ts">
	import { IconRotate } from '@tabler/icons-svelte';

	let { reset }: { reset: () => void } = $props();	

	let button: HTMLButtonElement;
	const animate = () => {
		button.animate([{ transform: 'rotate(0deg)' }, { transform: 'rotate(-360deg)' }], {
			duration: 500,
			iterations: 1,
			easing: 'ease-in-out'
		});
	};
</script>

<div class="tooltip tooltip-bottom" data-tip="Reset game">
	<button
		bind:this={button}
		onclick={() => {
			animate();
			reset();
		}}
	>
		<IconRotate class="swap-on h-10 w-10" />
	</button>
</div>
