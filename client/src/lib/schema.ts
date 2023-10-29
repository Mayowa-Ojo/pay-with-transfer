import z from "zod";

export const AccountRouteSchema = z.object({
   id: z.string().uuid(),
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
});

export type PaymentAccount = z.infer<typeof PaymentAccountSchema>;
