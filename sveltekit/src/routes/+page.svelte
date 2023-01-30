<script lang="ts">
	import Card from '../components/Card.svelte';
	import Hero from '../components/Hero.svelte';

	export let data: any = [];
	console.log();

	// // let Post = [];
	// // onMount(async () => {
	// // 	Post = await getPosts();
	// // });

	// const getPosts = async () => {
	// 	const res = await fetch('https://jsonplaceholder.typicode.com/posts');
	// 	const data = await res.json();
	// 	const filteredData = data.slice(0, 3);
	// 	return filteredData;
	// };
</script>

<Hero />
<!-- always use either this or by mount -->
<!-- {#await getPosts()}
	<p>.....Loading</p>
{:then data} -->
<!-- <div class="container">
	<div class="d-flex flex-row justify-content-between mx-4">
		{#each returnData.post as { title, url }}
			<div>
				<Card {title} {url} />
			</div>
		{/each}
	</div>
</div> -->

<div class="container">
	<div class="d-flex flex-row justify-content-between mx-4">
		{#each data.posts as { title, url }}
			<div>
				<Card {title} {url} />
			</div>
		{/each}
	</div>
</div>
<!-- {/await} -->

<!-- For avoiding expose api move from +page.js to +page.server -->
