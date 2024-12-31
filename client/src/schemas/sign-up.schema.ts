import { z } from "zod";

export const signUpSchema = z
  .object({
    name: z.string().nonempty("Name is required").max(100),
    email: z
      .string()
      .nonempty("Email is required")
      .max(100)
      .email("Please enter a valid email address"),
    password: z.string().nonempty("Password is required").min(8).max(100),
    passwordConfirmation: z
      .string()
      .nonempty("Password confirmation is required")
      .min(8)
      .max(100),
  })
  .refine((data) => data.passwordConfirmation === data.password, {
    message: "Passwords don't match",
    path: ["passwordConfirmation"],
  });

export type SignUpSchemaType = z.infer<typeof signUpSchema>;
