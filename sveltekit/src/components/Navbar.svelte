<script>
	import { page } from '$app/stores';
	$: console.log($page.route.id);
	$: routeId = $page.route.id;
	console.log(routeId);
	// will log the current route id to the console
</script>

<nav class="bg-dark text-white">
	<h4>Start Bootstrap</h4>
	<ul>
		<!-- can use above value for css like active/inactive.  -->
		<li><a href="/">Home</a></li>
		<li><a href="/about">About</a></li>
		<li><a href="/contact">Contact</a></li>
		<li><a href="/services">Services</a></li>
	</ul>
</nav>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Fira+Sans:wght@300&display=swap');
	nav {
		font-family: 'Fira Sans', sans-serif;
	}
	nav {
		display: flex;
		justify-content: space-between;
		background-color: #f8f9fa;
		padding: 0.6rem;
	}
	nav h4 {
		font-size: 1.3rem;
	}
	nav ul {
		display: inline-block;
		margin: 0;
		padding: 0;
	}
	nav li {
		display: inline-block;
		list-style: none;
	}
	nav a {
		text-decoration: none;
		padding: 0.5rem;
		color:#f8f9fa
	}
	nav a:hover {
		color: #007bff;
	}
</style>
