import z from "zod";

export const AccountRouteSchema = z.object({
   id: z.string().uuid(),
   txn: z.string().uuid(),
});

export const PaymentAccountSchema = z.object({
   id: z.string(),
   account_id: z.string(),
   payment_amount: z.number(),
   status: z.string(),
   expires_at: z.date(),
   created_at: z.date(),
   updated_at: z.date(),
   account_name: z.string(),
   account_number: z.string(),
   bank_name: z.string(),
   transaction_id: z.string(),
});

export type PaymentAccount = z.infer<typeof PaymentAccountSchema>;

export const TransactionSchema = z.object({
   id: z.string(),
   account_id: z.string(),
   ephemeral_account_id: z.string(),
   external_id: z.string(),
   amount: z.string(),
   currency: z.string(),
   account_name: z.string(),
   account_number: z.string(),
   bank_name: z.string(),
   status: z.string(),
   provider: z.string(),
   created_at: z.date(),
   updated_at: z.date(),
});

export type Transaction = z.infer<typeof TransactionSchema>;
