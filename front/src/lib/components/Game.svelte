<script lang="ts">
	import { getDailyCandidate } from '$lib/pocketBase';

	let { goalCandidate }: { goalCandidate: Awaited<ReturnType<typeof getDailyCandidate>> } = $props();
</script>

<main>
    <h1>Game</h1>
    <p>Goal: {goalCandidate.word} {goalCandidate.letter}</p>
</main>

