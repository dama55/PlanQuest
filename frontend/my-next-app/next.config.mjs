/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    swcMinify: true,
    webpack: (config, { isServer }) => {
        if (!isServer) {
            config.watchOptions = {
                poll: 500,         // 1秒ごとにファイル変更をポーリング
                aggregateTimeout: 300, // ファイル変更後、300ms待ってから再ビルド
            };
        }
        return config;
    },
};

export default nextConfig;
