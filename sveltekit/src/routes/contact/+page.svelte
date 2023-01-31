<script>
	import { applyAction, enhance } from '$app/forms';
	import { error } from '@sveltejs/kit';
	export let form;
	console.log(form);
</script>

{#if form?.success}
	<div class="alert alert-info" role="alert">
		{form.name}, your message has been sent. Thank you!
	</div>
{:else if form?.error}
	<div class="alert alert-danger" role="alert">{form.error}</div>
	<div class="container w-60 m-auto my-4">
		<form
			method="POST"
			use:enhance={() => {
				return async ({ result, error }) => {
					await applyAction(result, error);
				};
			}}
		>
			<div class="form-group">
				<label for="nameInput">Name</label>
				<input
					type="text"
					class="form-control"
					id="nameInput"
					placeholder="Enter your name"
					name="name"
					value={form?.name || ''}
					class:error={form?.errors?.name}
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
					value={form?.email || ''}
					class:error={form?.errors?.email}
				/>
			</div>
			<div class="form-group">
				<label for="messageInput">Message</label>
				<textarea
					class="form-control"
					value={form?.message || ''}
					id="messageInput"
					rows="3"
					name="message"
					class:error={form?.errors?.message}
				/>
			</div>
			<button type="submit" class="btn btn-primary my-4 w-100">Submit</button>
		</form>
	</div>
{:else}
	<div class="container w-60 m-auto my-4">
		<form method="POST" use:enhance>
			<div class="form-group">
				<label for="nameInput">Name</label>
				<input
					type="text"
					class="form-control"
					id="nameInput"
					placeholder="Enter your name"
					name="name"
					value={form?.name || ''}
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
					value={form?.email || ''}
				/>
			</div>
			<div class="form-group">
				<label for="messageInput">Message</label>
				<textarea
					class="form-control"
					value={form?.message || ''}
					id="messageInput"
					rows="3"
					name="message"
				/>
			</div>
			<button type="submit" class="btn btn-primary my-4 w-100">Submit</button>
		</form>
	</div>
{/if}
