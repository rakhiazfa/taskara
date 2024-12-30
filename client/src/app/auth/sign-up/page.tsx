import { Metadata } from "next";
import Link from "next/link";

import { SignUpForm } from "@/components/sign-up";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export const metadata: Metadata = {
  title: "Sign Up",
};

export default async function SignUp() {
  return (
    <main>
      <section className="app-container min-h-[500px] sm:min-h-[600px] flex justify-center items-center">
        <Card className="w-[400px]">
          <CardHeader>
            <CardTitle>Create new account</CardTitle>
          </CardHeader>
          <CardContent>
            <SignUpForm />
          </CardContent>
          <CardFooter>
            <div className="w-full flex flex-col gap-4">
              <p className="text-center">Already have an account?</p>
              <Button variant="secondary" className="w-full" asChild>
                <Link href="/auth/sign-in">Sign In</Link>
              </Button>
            </div>
          </CardFooter>
        </Card>
      </section>
    </main>
  );
}
