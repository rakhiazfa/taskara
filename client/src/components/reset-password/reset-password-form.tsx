"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form";

import { DEFAULT_ERROR_MESSAGE } from "@/constants/error-message.constant";
import { useToast } from "@/hooks/use-toast";
import {
  resetPasswordSchema,
  ResetPasswordSchemaType,
} from "@/schemas/reset-password.schema";

import { Button } from "../ui/button";
import { FormGroup } from "../ui/form-group";
import { Input } from "../ui/input";

export const ResetPasswordForm = () => {
  const router = useRouter();
  const { toast } = useToast();
  const searchParams = useSearchParams();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<ResetPasswordSchemaType>({
    resolver: zodResolver(resetPasswordSchema),
    defaultValues: {
      token: searchParams.get("token") ?? undefined,
    },
  });

  const [loading, setLoading] = useState(false);

  const onSubmit = async (payload: ResetPasswordSchemaType) => {
    try {
      setLoading(true);

      const response = await fetch("/api/auth/reset-password", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });
      const data = await response.json();

      if (!response.ok) throw new Error(data?.message);

      router.push("/auth/sign-in");
    } catch (error) {
      toast({
        variant: "destructive",
        title: "Error",
        description:
          error instanceof Error
            ? error.message
            : DEFAULT_ERROR_MESSAGE.UNEXPECTED_ERROR.MESSAGE,
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <form className="grid grid-cols-1 gap-4" onSubmit={handleSubmit(onSubmit)}>
      <FormGroup label="Password" error={errors.password?.message}>
        <Input
          {...register("password")}
          type="password"
          placeholder="Enter your password"
        />
      </FormGroup>
      <FormGroup
        label="Password Confirmation"
        error={errors.passwordConfirmation?.message}
      >
        <Input
          {...register("passwordConfirmation")}
          type="password"
          placeholder="Confirm your password"
        />
      </FormGroup>
      <Button type="submit" className="w-full mt-4" disabled={loading}>
        Continue
      </Button>
    </form>
  );
};
