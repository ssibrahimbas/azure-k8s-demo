import type { Response } from "@/types/response";
import axios from "axios";

type CounterService = {
  get(): Promise<Response<number>>;
  increment(): Promise<Response<number>>;
  decrement(): Promise<Response<number>>;
  set(value: number): Promise<Response<number>>;
  reset(): Promise<Response<number>>;
};

export const useCounterService = (): CounterService => ({
  get: async () => {
    const response = await axios.get(`/api/counter/v1`);
    if (response.status !== 200) {
      return {
        data: null,
        success: false,
        error: response.data.message ?? "Unknown error",
      };
    }
    return {
      data: response.data.data,
      success: true,
    };
  },
  increment: async () => {
    const response = await axios.post(`/api/counter/v1/inc`);
    if (response.status !== 200) {
      return {
        data: null,
        success: false,
        error: response.data.message ?? "Unknown error",
      };
    }
    return {
      data: response.data.data,
      success: true,
    };
  },
  decrement: async () => {
    const response = await axios.post(`/api/counter/v1/dec`);
    if (response.status !== 200) {
      return {
        data: null,
        success: false,
        error: response.data.message ?? "Unknown error",
      };
    }
    return {
      data: response.data.data,
      success: true,
    };
  },
  set: async (count: number) => {
    const response = await axios.post(`/api/counter/v1/set`, {
      count,
    });
    if (response.status !== 200) {
      return {
        data: null,
        success: false,
        error: response.data.message ?? "Unknown error",
      };
    }
    return {
      data: response.data.data,
      success: true,
    };
  },
  reset: async () => {
    const response = await axios.post(`/api/counter/v1/reset`);
    if (response.status !== 200) {
      return {
        data: null,
        success: false,
        error: response.data.message ?? "Unknown error",
      };
    }
    return {
      data: response.data.data,
      success: true,
    };
  },
});
