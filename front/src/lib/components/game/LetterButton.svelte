<script lang="ts">
	let {
		letter,
		onClick,
		isCenter,
		bonusLetter
	}: { letter: string; onClick: (letter: string) => void; isCenter: boolean; bonusLetter: string } =
		$props();

	const isBonus = letter === bonusLetter;
	const color = isCenter ? 'primary' : isBonus ? 'secondary' : 'neutral';
</script>

<button
	class="link text-neutral hover:border-primary hover:text-primary-content relative h-20 w-20 font-bold"
	style="--path-bg-base: var(--color-{color}-content); --path-bg-hover: var(--color-{color})"
	onclick={() => onClick(letter)}
>
	<div class="absolute top-0 left-0 z-10 flex h-full w-full items-center justify-center text-5xl">
		{letter}
	</div>
	<div class="absolute top-0 left-0 h-full w-full">
		<svg
			xmlns="http://www.w3.org/2000/svg"
			height="100%"
			width="100%"
			viewBox="-15.51 -52.47 61.01 52.97"
			vector-effect="non-scaling-stroke"
		>
			<path
				id="letter"
				fill="none"
				stroke-width="2"
				stroke-linejoin="round"
				stroke-linecap="round"
				stroke-dasharray="none"
				d="M 0,0 L 30,0 L 45,-25.98 L 30,-51.96 L 0,-51.96 L -15,-25.98 L 0,0 "
			>
			</path>
		</svg>
	</div>
	<div class="absolute top-[-8px] left-[-8px] h-full w-full">
		<svg
		xmlns="http://www.w3.org/2000/svg"
		height="120%"
		width="120%"
		viewBox="-15.51 -52.47 61.01 52.97"
		vector-effect="non-scaling-stroke"
	>
			<path
				fill="none"
				stroke-width="var(--focus-stroke-width)"
				stroke-linejoin="round"
				stroke-linecap="round"
				stroke="var(--path-bg-hover)"
				d="M 0,0 L 30,0 L 45,-25.98 L 30,-51.96 L 0,-51.96 L -15,-25.98 L 0,0 "
			>
			</path>
	</svg>
	</div>
</button>

<style>
	button {
		transition-property: color, background-color, border-color, box-shadow;
		transition-timing-function: cubic-bezier(0, 0, 0.2, 1);
		transition-duration: 0.2s;
		--focus-stroke-width: 0px;
		&:focus-visible {
			outline-width: 0px;
			--focus-stroke-width: 2px;
		}
		&:active:not(.btn-active) {
			translate: 0 0.5px;
			--btn-bg: color-mix(in oklab, var(--btn-color, var(--color-base-200) /* var(--color-base-200) */), #000 5%);
			--btn-border: color-mix(in oklab, var(--btn-color, var(--color-base-200) /* var(--color-base-200) */), #000 7%);
			--btn-shadow: 0 0 0 0 oklch(0% 0 0/0), 0 0 0 0 oklch(0% 0 0/0);
		}
	}
	button #letter {
		fill: var(--path-bg-base);
		stroke: var(--path-bg-hover);
	}
	button:hover #letter {
		fill: var(--path-bg-hover);
		stroke: var(--path-bg-base);
	}
</style>
