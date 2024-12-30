import { Metadata } from "next";

import { ResetPasswordForm } from "@/components/reset-password";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export const metadata: Metadata = {
  title: "Reset Password",
};

export default async function ResetPassword() {
  return (
    <main>
      <section className="app-container min-h-[500px] sm:min-h-[600px] flex justify-center items-center py-6">
        <Card className="w-[400px]">
          <CardHeader>
            <CardTitle>Enter your new password</CardTitle>
            <CardDescription>
              Your new password must be different from your previous password.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <ResetPasswordForm />
          </CardContent>
        </Card>
      </section>
    </main>
  );
}
