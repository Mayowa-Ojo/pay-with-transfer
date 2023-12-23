import { useQuery } from "@tanstack/react-query";
import { useHttpRequest } from "./http";
import constants from "./constants";
import { PaymentAccount, Transaction } from "./schema";

const useFetchPaymentAccount = (id: string) => {
   const httpRequest = useHttpRequest<PaymentAccount>();

   return useQuery(
      [constants.QUERY_KEY_FETCH_PAYMENT_ACCOUNT],
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

const useFetchTransaction = (id: string, isPolling: boolean) => {
   const httpRequest = useHttpRequest<Transaction>();

   return useQuery(
      [constants.QUERY_KEY_FETCH_TRANSACTION],
      () => {
         return httpRequest(constants.API_PATH_FETCH_TRANSACTION(id), "get");
      },
      {
         enabled: !!id && isPolling,
         onSuccess: (resp) => {
            console.log(resp);
         },
         refetchInterval: 8 * 1000,
      }
   );
};

export const query = {
   fetchPaymentAccount: useFetchPaymentAccount,
   fetchTransaction: useFetchTransaction,
};
