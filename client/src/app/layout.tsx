import "../assets/css/globals.css";

import type { Metadata } from "next";
import { Inter } from "next/font/google";

import { Footer } from "@/components/ui/footer";

import Providers from "./providers";

const inter = Inter({
  variable: "--font-inter",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: {
    default: `${process.env.NEXT_PUBLIC_APP_NAME}`,
    template: `%s | ${process.env.NEXT_PUBLIC_APP_NAME}`,
  },
  description: `${process.env.NEXT_PUBLIC_APP_DESCRIPTION}`,
  keywords: ["Project Management Tool", "Task Management Tool"],
  authors: [{ name: `${process.env.NEXT_PUBLIC_APP_AUTHOR}` }],
  creator: `${process.env.NEXT_PUBLIC_APP_AUTHOR}`,
  publisher: `${process.env.NEXT_PUBLIC_APP_AUTHOR}`,
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.variable}`}>
        <Providers>{children}</Providers>
        <Footer />
      </body>
    </html>
  );
}
