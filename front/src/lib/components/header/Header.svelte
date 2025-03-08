<script lang="ts">
	import Return from '$lib/components/header/Return.svelte';
	import ResetButton from '$lib/components/header/ResetButton.svelte';
	import Home from '$lib/components/header/Home.svelte';
	import { onMount } from 'svelte';
	import { themeChange } from 'theme-change';
	import { IconSun, IconMoon } from '@tabler/icons-svelte';

	// NOTE: the element that is using one of the theme attributes must be in the DOM on mount
	onMount(() => {
		themeChange(false);
		// 👆 false parameter is required for svelte
	});

	let { reset, onlyTheme }: { reset?: () => void; onlyTheme?: boolean } = $props();
</script>

{#if !onlyTheme}
	<div class="bg-base-200 bg-opacity-80 fixed top-0 left-0 z-10 m-2 h-[56px] rounded-xl p-2">
		<Return />
		<Home />
		{#if reset}
			<ResetButton {reset} />
		{/if}
	</div>
{/if}
<div class="bg-base-200 bg-opacity-80 fixed top-0 right-0 z-10 m-2 h-[56px] rounded-xl p-2">
	<div class="tooltip tooltip-bottom before:whitespace-pre-wrap">
		<div class="tooltip-content max-w-17">Changer de thème</div>
		<label class="swap swap-rotate">
			<!-- this hidden checkbox controls the state -->
			<input type="checkbox" data-toggle-theme="dark,light" data-act-class="ACTIVECLASS" hidden />
			<IconSun class="swap-on h-10 w-10" stroke={2} />
			<IconMoon class="swap-off h-10 w-10" stroke={2} />
		</label>
	</div>
</div>
<div class="mb-[72px] w-full sm:m-2"></div>
