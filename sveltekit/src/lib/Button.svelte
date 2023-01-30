<script lang="ts">
	export let size;
	export let clas;
	export let id;

	console.log(size, clas, id);
</script>

<button class="btn btn-success">
	<slot />
</button>

<style>
</style>
