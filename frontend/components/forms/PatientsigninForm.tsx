"use client"
 
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import {Form} from "@/components/ui/form"
import CustomFormField from "../CustomFormField"
import "react-phone-number-input/style.css";
import SubmitButton from "../SubmitButton"
import { useState } from "react"
import { UserFormValidation } from "./validation"
import { useRouter } from "next/navigation";
import { loginUser } from "@/app/api/register/register"

export enum FormFieldType {
  INPUT = "input",
  TEXTAREA = "textarea",
  PHONE_INPUT = "phoneInput",
  CHECKBOX = "checkbox",
  DATE_PICKER = "datePicker",
  SELECT = "select",
  SKELETON = "skeleton",
  PASSWORD= "password",
}

export interface user{
  name: string;
  email: string;
  phone: string;
  password: string;
  role: string;
}

 
const PatientsigninForm =()=> {
  const [isLoading, setIsLoading] = useState(false)
  const router = useRouter();
  // 1. Define form.
  const form = useForm<z.infer<typeof UserFormValidation>>({
    resolver: zodResolver(UserFormValidation),
    defaultValues: {
      name: "",
      email: "",
      phone: "",
      password: "",
    },
  })
 
  const onSubmit = async (values: z.infer<typeof UserFormValidation>) => {
    setIsLoading(true);

    try {
      const user:user = {
        name: values.name,
        email: values.email,
        phone: values.phone,
        password: values.password,
        role: "Patient",
      };
      
      const response = await loginUser(user);
 
      if (response.status === 201) {
        router.push(`/patients/${response.data.user_id}/register`);
      }else{
        router.push(`/patients/${response.data.user_id}/dashboard`);
      }
      
    } catch (error) {
      console.log(error);
    }

    setIsLoading(false);
  };
  return (
    <Form {...form}>
    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6 flex-1">
      <section className="mb-12 space-y-4">
        <h1 className="header">Hi there 👋</h1>
        <p className="text-dark-700">Get started with appointments.</p>
      </section>
        <CustomFormField
          fieldType={FormFieldType.INPUT}
          control={form.control}
          name="name"
          label="Full name"
          placeholder="John Doe"
          iconSrc="/assets/icons/user.svg"
          iconAlt="user"
        />

        <CustomFormField
          fieldType={FormFieldType.INPUT}
          control={form.control}
          name="email"
          label="Email"
          placeholder="johndoe@gmail.com"
          iconSrc="/assets/icons/email.svg"
          iconAlt="email"
        />

        <CustomFormField
          fieldType={FormFieldType.PHONE_INPUT}
          control={form.control}
          name="phone"
          label="Phone number"
          placeholder="(555) 123-4567"
        />

        <CustomFormField
          fieldType={FormFieldType.PASSWORD}
          control={form.control}
          name="password"
          label="password"
          placeholder="********"
          iconSrc="/assets/icons/password.svg"
          iconAlt="password"
        />
       <SubmitButton isLoading={isLoading}>Get Started</SubmitButton>
    </form>
  </Form>
  )
}

export default PatientsigninForm