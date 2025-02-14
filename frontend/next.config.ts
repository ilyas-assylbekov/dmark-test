import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'export',
  distDir: 'dist',
  images: {
    unoptimized: true
  },
  assetPrefix: '.',
  // Add this to handle static export
  trailingSlash: true,
};

export default nextConfig;
