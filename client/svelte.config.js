import adapter from '@sveltejs/adapter-static';
import path from 'path';

export default {
    kit: {
        adapter: adapter({
            pages: path.resolve('../server/static'),
            assets: path.resolve('../server/static'),
            fallback: 'index.html', // or 'index.html' depending on your setup
        })
    },
    compilerOptions: {
        runes: true,
    },

};
