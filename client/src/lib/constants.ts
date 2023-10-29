const constants = {
   //[app]
   API_BASE_URL: import.meta.env.VITE_APP_API_BASE_URL,
   //[api-paths]
   API_PATH_GENERATE_PAYMENT_ACCOUNT: "api/v1/accounts/ephemeral",
   API_PATH_FETCH_PAYMENT_ACCOUNT: (id: string) =>
      `api/v1/accounts/${id}/ephemeral`,
   //[query-keys]
   QUERY_KEY_ETCH_PAYMENT_ACCOUNT: "payment-account",
   //[status-codes]
   HTTP_STATUS_UNAUTHORIZED: 401,
   HTTP_STATUS_FORBIDEN: 403,
} as const;

export default constants;
