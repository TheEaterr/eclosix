<script lang="ts">
	import { getDailyCandidate } from '$lib/pocketBase';

	export let goalCandidate: Awaited<ReturnType<typeof getDailyCandidate>>;
</script>

<main>
    <h1>Game</h1>
    <p>Goal: {goalCandidate.word} {goalCandidate.letter}</p>
</main>

