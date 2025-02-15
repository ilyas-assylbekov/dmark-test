import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'export',
  distDir: 'dist',
  images: {
    unoptimized: true
  },
  assetPrefix: '.',
  // Это для обработки статического экспорта
  trailingSlash: true,
};

export default nextConfig;
