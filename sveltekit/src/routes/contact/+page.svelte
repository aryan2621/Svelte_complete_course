<script>
	export let form;
	console.log(form);
</script>

<div class="container w-60 m-auto my-4">
	<form method="POST">
		<div class="form-group">
			<label for="nameInput">Name</label>
			<input
				type="text"
				class="form-control"
				id="nameInput"
				placeholder="Enter your name"
				name="name"
			/>
		</div>
		<div class="form-group">
			<label for="emailInput">Email address</label>
			<input
				type="email"
				class="form-control"
				id="emailInput"
				placeholder="Enter your email"
				name="email"
			/>
		</div>
		<div class="form-group">
			<label for="messageInput">Message</label>
			<textarea class="form-control" id="messageInput" rows="3" name="message" />
		</div>
		<button type="submit" class="btn btn-primary my-4 w-100">Submit</button>
	</form>
</div>
