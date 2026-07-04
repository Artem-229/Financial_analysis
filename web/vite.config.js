import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            '/login': 'http://localhost:8080',
            '/registration': 'http://localhost:8080',
            '/app': 'http://localhost:8080',
        },
    },
});
