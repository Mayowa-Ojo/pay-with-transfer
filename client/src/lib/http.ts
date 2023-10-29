import ky from "ky";
import constants from "./constants";

type HttpMethod = "get" | "post" | "put" | "patch" | "head" | "delete";

export interface APIResponse<T> {
   code: "00" | "01" | "02";
   message: string;
   data: T;
}

export const useHttpRequest = <T>() => {
   return async (
      path: string,
      method: HttpMethod,
      headers: Record<string, string> = {},
      payload?: unknown,
      params?: URLSearchParams
   ) => {
      return (await ky[method](path, {
         prefixUrl: constants.API_BASE_URL,
         searchParams: params,
         json: payload || undefined,
         headers: {
            "Content-Type": "application/json",
            ...headers,
         },
      }).json()) as APIResponse<T>;
   };
};
