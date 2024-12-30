import Link from "next/link";

import { AppLogo } from "./app-logo";

export const Footer = () => {
  return (
    <footer className="w-full bg-secondary py-6">
      <div className="app-container">
        <AppLogo className="mb-4" />
        <div className="flex flex-col sm:flex-row justify-between sm:items-center gap-4">
          <p className="text-muted-foreground order-2 sm:order-1">
            Copyright &copy; 2024 {process.env.NEXT_PUBLIC_APP_NAME}
          </p>
          <div className="flex items-center gap-6 order-1 sm:order-2">
            <Link
              href="/privacy-policy"
              className="text-muted-foreground hover:text-neutral-600"
            >
              Privacy Policy
            </Link>
            <Link
              href="/terms"
              className="text-muted-foreground hover:text-neutral-600"
            >
              Terms of Service
            </Link>
          </div>
        </div>
      </div>
    </footer>
  );
};
