import { NextResponse } from "next/server";

export async function loginUser(request: {
    name: string;
    email: string;
    phone: string;
    password: string;
    role: string;
}) {
    const { name, email, phone, password, role } = request;

    const response = await fetch("http://localhost:8080/v1/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, phone, password, role }),
        });
    
    if (response.ok) {
        return response.json();
    }

    return NextResponse.error();
}

export async function test(){
    console.log("test")
    const response = await fetch("http://localhost:8080/v1/auth/test",{
        method:"GET",
        headers: { "Content-Type": "application/json" },
    })
    return response.json()
}
 