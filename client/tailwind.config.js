/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      './src/**/*.{js,svelte}',
    ],
    theme: {
        extend: {
            colors: {
                'primary': '#E51636',
                'error': '#750c05',
                'gray': {
                    100: '#EEEEEE',
                    500: '#BBBBBB',
                    900: '#333333',
                }
            },
        },
    },
    plugins: [],
  };