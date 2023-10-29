import { useMutation } from "@tanstack/react-query";
import { APIResponse, useHttpRequest } from "./http";
import constants from "./constants";
import { useNavigate } from "@tanstack/react-router";
import { PaymentAccount } from "./schema";

export interface GeneratePaymentAccountPayload {
   amount: number;
   session_id: string;
}

const useGeneratePaymentAccount = () => {
   const httpRequest = useHttpRequest<PaymentAccount>();
   const navigate = useNavigate();

   return useMutation<
      APIResponse<unknown>,
      unknown,
      GeneratePaymentAccountPayload,
      unknown
   >(
      (payload) => {
         return httpRequest(
            constants.API_PATH_GENERATE_PAYMENT_ACCOUNT,
            "post",
            {},
            payload
         );
      },
      {
         onSuccess: async (resp) => {
            console.log(resp);
            const data = resp.data;
            navigate({
               to: "/payment/account",
               search: { id: (data as PaymentAccount).id },
            });
         },
      }
   );
};

export const mutation = {
   generatePaymentAccount: useGeneratePaymentAccount,
};
