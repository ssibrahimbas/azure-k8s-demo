export type AppConfig = {
  apiBaseUrl: string;
};

export const useConfig = (): AppConfig => ({
  apiBaseUrl: import.meta.env.API_BASE_URL,
});
