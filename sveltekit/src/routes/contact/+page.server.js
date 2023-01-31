import { fail } from '@sveltejs/kit';
import { object, string } from 'yup';

const schema = object({
	name: string().min(2).required('Please enter your name'),
	email: string().email('Please enter a valid email').required('Please enter your email'),
	message: string().required('Please enter your message')
});

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
		if (!message) {
			// return {
			// 	error: 'Please enter your name'
			// };
			// return fail(400, { message: 'Please enter your name', email, name, status: 'error' });
			// return fail(400, { error: 'Please enter your message', email, name, status: 'error' });
			// to make value be there and not disappear and only
			// missing vaule to be highlighted
		}
		if (!email) {
			// return {
			// 	error: 'Please enter your name'
			// };
			// return fail(400, { error: 'Please enter your email', message, name, status: 'error' });
			// to make value be there and not disappear and only
			// missing vaule to be highlighted
		}
		if (!name) {
			// return {
			// 	error: 'Please enter your name'
			// };
			// return fail(400, { error: 'Please enter your name', email, message, status: 'error' });
			// to make value be there and not disappear and only
			// missing vaule to be highlighted
		}
		try {
			const result = await schema.validate({ name, email, message }, { abortEarly: false });
			console.log('result', { result });
			return {
				success: true,
				name,
				email,
				message
			};
		} catch (error) {
			// console.log('error', { error });
			const errors = error.inner.reduce((acc, curr) => {
				acc[curr.path] = curr.message;
				return acc;
			}, {});
			console.log('errors', { errors });
			return {
				errors,
				name,
				email,
				message
			};
		}
	}
};
