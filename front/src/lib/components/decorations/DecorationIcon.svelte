<script lang="ts">
	import {
		IconWriting,
		IconSignature,
		IconWorld,
		IconBallpen,
		IconLetterCaseToggle,
		IconBook2,
		IconPencil,
		IconFile
	} from '@tabler/icons-svelte';
	import { DecorationPossibilities } from './decorationPossibilities';

	export let decoration: DecorationPossibilities;
</script>

{#if decoration === DecorationPossibilities.WRITING}
	<IconWriting />
{:else if decoration === DecorationPossibilities.SIGNATURE}
	<IconSignature />
{:else if decoration === DecorationPossibilities.WORLD}
	<IconWorld />
{:else if decoration === DecorationPossibilities.BALLPEN}
	<IconBallpen />
{:else if decoration === DecorationPossibilities.LETTER_CASE_TOGGLE}
	<IconLetterCaseToggle />
{:else if decoration === DecorationPossibilities.BOOK}
	<IconBook2 />
{:else if decoration === DecorationPossibilities.PENCIL}
	<IconPencil />
{:else if decoration === DecorationPossibilities.FILE}
	<IconFile />
{:else}
	<p>Unknown decoration</p>
{/if}

<style>
</style>
