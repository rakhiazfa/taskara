import { AxiosError } from "axios";
import { NextRequest, NextResponse } from "next/server";

import { DEFAULT_ERROR_MESSAGE } from "@/constants/error-message.constant";
import api from "@/lib/api";

export async function POST(request: NextRequest) {
  try {
    const { name, email, password, passwordConfirmation } =
      await request.json();

    const { status, data } = await api.post("/auth/sign-up", {
      name,
      email,
      password,
      passwordConfirmation,
    });

    return NextResponse.json(
      data ?? { message: "Successfully create a new account" },
      { status },
    );
  } catch (error) {
    if (error instanceof AxiosError)
      return NextResponse.json(
        error.response?.data ?? {
          message: DEFAULT_ERROR_MESSAGE.UNEXPECTED_ERROR.MESSAGE,
        },
        { status: error.status || 500 },
      );

    return NextResponse.json(
      { message: DEFAULT_ERROR_MESSAGE.UNEXPECTED_ERROR.MESSAGE },
      { status: 500 },
    );
  }
}
