<script lang="ts">
	let {
		letter,
		onClick,
		isCenter,
		bonusLetter
	}: { letter: string; onClick: (letter: string) => void; isCenter: boolean; bonusLetter: string } =
		$props();

	const isBonus = $derived(letter === bonusLetter);
	const color = $derived(isCenter ? 'primary' : isBonus ? 'secondary' : 'neutral');
	const classes = {
		primary:
			'text-primary hover:text-primary-content fill-(--color-primary-content) hover:fill-(--color-primary) stroke-primary hover:stroke-primary-content',
		secondary:
			'text-secondary hover:text-secondary-content fill-(--color-secondary-content) hover:fill-(--color-secondary) stroke-secondary hover:stroke-secondary-content',
		neutral:
			'text-neutral hover:text-neutral-content fill-(--color-neutral-content) hover:fill-(--color-neutral) stroke-neutral hover:stroke-neutral-content'
	};
</script>

<button
	class="link text-neutral {classes[color]} relative h-22 w-22 m-1.5 font-bold"
	type="button"
	onclick={() => onClick(letter)}
>
	<div class="absolute top-0 left-0 z-1 flex h-full w-full items-center justify-center text-5xl">
		{letter.toUpperCase()}
	</div>
	<div class="absolute top-[-6.6px] left-[-6.6px] h-full w-full -z-10" >
		<svg
			xmlns="http://www.w3.org/2000/svg"
			height="115%"
			width="115%"
			viewBox="-15.51 -52.47 61.01 52.97"
			vector-effect="non-scaling-stroke"
		>
			<path
				id="letter"
				stroke-width="2"
				stroke-linejoin="round"
				stroke-linecap="round"
				stroke-dasharray="none"
				d="M 0,0 L 30,0 L 45,-25.98 L 30,-51.96 L 0,-51.96 L -15,-25.98 L 0,0 "
			>
			</path>
		</svg>
	</div>
	<div class="absolute top-[-13.2px] left-[-13.2px] h-full w-full -z-10">
		<svg
			xmlns="http://www.w3.org/2000/svg"
			height="150%"
			width="150%"
			viewBox="-15.51 -52.47 61.01 52.97"
			vector-effect="non-scaling-stroke"
		>
			<path
				fill="none"
				stroke-width="var(--focus-stroke-width)"
				stroke-linejoin="round"
				stroke-linecap="round"
				stroke="var(--color-{color})"
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
		}
	}
</style>
