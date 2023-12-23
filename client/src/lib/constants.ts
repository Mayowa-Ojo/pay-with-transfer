const constants = {
   //[app]
   API_BASE_URL: import.meta.env.VITE_APP_API_BASE_URL,
   //[api-paths]
   API_PATH_GENERATE_PAYMENT_ACCOUNT: "api/v1/accounts/ephemeral",
   API_PATH_FETCH_PAYMENT_ACCOUNT: (id: string) =>
      `api/v1/accounts/${id}/ephemeral`,
   API_PATH_FETCH_TRANSACTION: (id: string) => `api/v1/transactions/${id}`,
   //[query-keys]
   QUERY_KEY_FETCH_PAYMENT_ACCOUNT: "payment-account",
   QUERY_KEY_FETCH_TRANSACTION: "transaction",
   //[status-codes]
   HTTP_STATUS_UNAUTHORIZED: 401,
   HTTP_STATUS_FORBIDEN: 403,
} as const;

export default constants;
