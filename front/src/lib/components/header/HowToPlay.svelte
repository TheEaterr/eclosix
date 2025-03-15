<script lang="ts">
	import { IconHelp } from '@tabler/icons-svelte';

	let button: HTMLButtonElement;
	let helpModal: HTMLDialogElement;
	const animate = () => {
		button?.animate(
			[{ transform: 'scale(1)' }, { transform: 'scale(0.9)' }, { transform: 'scale(1)' }],
			{
				duration: 500,
				iterations: 1,
				easing: 'ease-in-out'
			}
		);
	};
</script>

<dialog id="tutorial-modal" class="modal" bind:this={helpModal}>
	<div class="modal-box bg-base-200 w-11/12 max-w-5xl">
		<h3 class="small-title text-primary text-3xl font-bold">Comment jouer à Éclosix</h3>
		<hr class="bg-neutral h-[1px] rounded" />
		<div class="text-justify">
			<p class="mt-2 py-2">
				Le but d'Éclosix est d'écrire 12 mots avec les lettres disponibles pour obtenir un maximum
				de points.
			</p>
			<p class="py-2">
				Vous pouvez utilisez autant de fois chaque lettre, mais chaque mot doit faire au moins 4
				lettres et contenir la lettre centrale. Les noms propres et noms composés sont exclus de la
				liste.
			</p>
			<p class="mb-2 py-2">
				Un mot de 4 lettres vaut 2 points, un mot de 5 lettres vaut 4 points, un mot de 6 lettres
				vaut 6 points et un mot de 7 lettres vaut 12 points. Chaque lettre supplémentaire rapporte 3
				points de plus. La lettre en bleu rapporte 5 points supplémentaire si elle est utilisée dans
				le mot. Enfin, 7 points supplémentaires sont accordés si le mot utilise toutes les lettres.
			</p>
			<p class="mb-2 py-2">Il existe au moins un mot utilisant toutes les lettres par problème.</p>
			<p class="mb-2 py-2">Un nouveau puzzle est généré chaque jour.</p>
			<hr class="bg-neutral h-[1px] rounded" />
			<div class="modal-action">
				<a
					class="btn btn-neutral-special text-lg"
					onclick={() => {
						helpModal?.close();
					}}
					href={null}
				>
					Fermer
				</a>
			</div>
		</div>
	</div>
</dialog>
<div class="tooltip tooltip-bottom" data-tip="Aide">
	<button
		class="link"
		bind:this={button}
		onclick={() => {
			animate();
			helpModal?.showModal();
		}}
	>
		<IconHelp class="swap-on h-10 w-10" />
	</button>
</div>
