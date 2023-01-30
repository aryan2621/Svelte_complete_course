export const actions = {
	default: async ({ request }) => {
		console.log('Contact form submitted!');
		const formData = await request.formData();
		const name = formData.get('name');
		const email = formData.get('email');
		const message = formData.get('message');

		console.log('Name:', name);
		console.log('Email:', email);
		console.log('Message:', message);
		return {
			message: 'Thank you for your message!'
		};
	}
};
