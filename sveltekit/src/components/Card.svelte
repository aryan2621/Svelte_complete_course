<script lang="ts">
	export let title: string;
	export let url: string;

	console.log(title, url);
</script>

<main>
	<div class="card" style="width: 18rem;">
		<!-- svelte-ignore a11y-img-redundant-alt -->
		<img class="card-img-top" src="https://picsum.photos/200" alt="Card image cap" />
		<div class="card-body">
			<p class="card-text">{title}</p>
		</div>
	</div>
</main>
