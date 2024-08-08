import { z } from "zod";


export const UserFormValidation = z.object({
    name: z
      .string()
      .min(2, "Name must be at least 2 characters")
      .max(50, "Name must be at most 50 characters"),
    email: z.string().email("Invalid email address"),
    phone: z
      .string()
      .refine((phone) => /^\+\d{10,15}$/.test(phone), "Invalid phone number"),
    password: z
      .string()
      .refine((password) => /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/.test(password), "Password must be  minimum of 8 characters, at least one uppercase letter, one lowercase letter, one number and one special character"),
  });