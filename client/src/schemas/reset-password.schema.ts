import { z } from "zod";

export const resetPasswordSchema = z
  .object({
    token: z.string().nonempty("Token is required").max(100),
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

export type ResetPasswordSchemaType = z.infer<typeof resetPasswordSchema>;
