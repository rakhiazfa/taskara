"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form";

import { DEFAULT_ERROR_MESSAGE } from "@/constants/error-message.constant";
import { useToast } from "@/hooks/use-toast";
import { signUpSchema, SignUpSchemaType } from "@/schemas/sign-up.schema";

import { Button } from "../ui/button";
import { FormGroup } from "../ui/form-group";
import { Input } from "../ui/input";

export const SignUpForm = () => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm<SignUpSchemaType>({
    resolver: zodResolver(signUpSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      passwordConfirmation: "",
    },
  });

  const [loading, setLoading] = useState(false);

  const onSubmit = async (payload: SignUpSchemaType) => {
    try {
      setLoading(true);

      const response = await fetch("/api/auth/sign-up", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });
      const data = await response.json();

      if (!response.ok) {
        switch (response.status) {
          case 409:
            setError(data?.field, { type: "validate", message: data?.message });
            break;
          default:
            throw new Error(data?.message);
        }
        return;
      }

      router.push("/dashboard");
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
      <FormGroup label="Name" error={errors.name?.message}>
        <Input
          {...register("name")}
          type="text"
          placeholder="Enter your name"
        />
      </FormGroup>
      <FormGroup label="Email" error={errors.email?.message}>
        <Input
          {...register("email")}
          type="text"
          placeholder="Enter your email"
        />
      </FormGroup>
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
