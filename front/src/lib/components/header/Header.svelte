<script lang="ts">
	import Return from '$lib/components/header/Return.svelte';
	import ResetButton from '$lib/components/header/ResetButton.svelte';
	import Home from '$lib/components/header/Home.svelte';

	let { reset }: { reset?: () => void; } = $props();
</script>

<div class="bg-base-200 bg-opacity-80 fixed top-0 left-0 z-10 m-2 h-[56px] rounded-xl p-2">
	<Return />
	<Home />
	{#if reset}
		<ResetButton {reset} />
	{/if}
</div>
<div class="mb-[72px] w-full sm:m-2"></div>
