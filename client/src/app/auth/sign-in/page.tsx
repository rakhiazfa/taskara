import { Metadata } from "next";
import Link from "next/link";
import { FcGoogle } from "react-icons/fc";

import { SignInForm } from "@/components/sign-in/sign-in-form";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export const metadata: Metadata = {
  title: "Sign In",
};

export default async function SignIn() {
  return (
    <main>
      <section className="app-container min-h-[500px] sm:min-h-[600px] flex justify-center items-center py-6">
        <Card className="w-[400px]">
          <CardHeader>
            <CardTitle>Sign In</CardTitle>
          </CardHeader>
          <CardContent>
            <SignInForm />
          </CardContent>
          <CardFooter>
            <div className="w-full flex flex-col gap-4">
              <Button variant="outline" className="w-full" asChild>
                <Link
                  href={`${process.env.NEXT_PUBLIC_BACKEND_URL}/oauth/sign-in/google`}
                >
                  <FcGoogle />
                  Continue With Google
                </Link>
              </Button>
              <p className="text-center">Don&apos;t have an account?</p>
              <Button variant="secondary" className="w-full" asChild>
                <Link href="/auth/sign-up">Sign Up</Link>
              </Button>
            </div>
          </CardFooter>
        </Card>
      </section>
    </main>
  );
}
