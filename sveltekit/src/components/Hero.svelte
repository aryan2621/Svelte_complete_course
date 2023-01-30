<script>
	import Button from '../lib/Button.svelte';
</script>

<div class="container  p-4">
	<div class="jumbotron">
		<h1 class="display-4">Hello, world!</h1>
		<p class="lead">
			This is a simple hero unit, a simple jumbotron-style component for calling extra attention to
			featured content or information.
		</p>
		<hr class="my-4" />
		<p>
			It uses utility classes for typography and spacing to space content out within the larger
			container.
		</p>
		<p class="lead">
			<!-- passign argument -->
			<Button clas="secondary" id="learn-more" size="lg">Learn more</Button>
			<!-- <a class="btn btn-primary btn-lg" href="/" role="button">Learn more</a> -->
		</p>
	</div>
</div>
<div class="container">
	<div class="card">
		<div class="card-header">Quote</div>
		<div class="card-body">
			<blockquote class="blockquote mb-0">
				<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer posuere erat a ante.</p>
				<footer class="blockquote-footer">
					Someone famous in <cite title="Source Title">Source Title</cite>
				</footer>
			</blockquote>
		</div>
	</div>
</div>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Fira+Sans:wght@300&display=swap');
	.container {
		padding: 10px;
		font-family: 'Fira Sans', sans-serif;
	}
</style>
