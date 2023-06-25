import Drawer from "../components/Drawer/drawer";
import "./globals.css";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Library System",
  description:
    "Discover and manage your favorite books with our online library platform. Easily search, borrow, and return books! It's also easy to manage users.",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <Drawer />
        {children}
      </body>
    </html>
  );
}
