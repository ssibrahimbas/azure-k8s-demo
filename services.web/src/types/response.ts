export type Response<T> = {
  data: T | null;
  success: boolean;
  error?: string;
};
