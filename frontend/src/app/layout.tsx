// Корневой макет приложения
import type { Metadata } from "next";
import "./globals.css";

// Метаданные приложения для SEO и заголовка страницы
export const metadata: Metadata = {
  title: "DMARK Task Manager",
  description: "Task Management Application",
};

// Корневой компонент макета
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className="antialiased">
        {children}
      </body>
    </html>
  );
}
