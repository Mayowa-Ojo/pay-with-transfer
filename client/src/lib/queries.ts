import { useQuery } from "@tanstack/react-query";
import { useHttpRequest } from "./http";
import constants from "./constants";
import { PaymentAccount } from "./schema";

const useFetchPaymentAccount = (id: string) => {
   const httpRequest = useHttpRequest<PaymentAccount>();

   return useQuery(
      [constants.QUERY_KEY_ETCH_PAYMENT_ACCOUNT],
      () => {
         return httpRequest(
            constants.API_PATH_FETCH_PAYMENT_ACCOUNT(id),
            "get"
         );
      },
      {
         enabled: !!id,
         onSuccess: (resp) => {
            console.log(resp);
         },
      }
   );
};

export const query = {
   fetchPaymentAccount: useFetchPaymentAccount,
};
