import React from "react";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import {
   Navigate,
   Outlet,
   Route,
   Router,
   RouterContext,
   RouterProvider,
   useNavigate,
   useRouter,
   useSearch,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
   faArrowLeftLong,
   faArrowRight,
   faBuildingColumns,
   faCheckCircle,
   faClone,
   faExclamationTriangle,
   faInfoCircle,
   faX,
} from "@fortawesome/free-solid-svg-icons";
import { useCopyToClipboard, useInterval, useTimeout } from "usehooks-ts";
import { mutation } from "./lib/mutations";
import { nanoid } from "nanoid";
import { query } from "./lib/queries";
import { AccountRouteSchema } from "./lib/schema";

dayjs.extend(utc);

const queryClient = new QueryClient();

const routerContext = new RouterContext<{
   queryClient: typeof queryClient;
}>();

const rootRoute = routerContext.createRootRoute({
   component: () => {
      return (
         <main className="w-full h-screen flex items-center justify-center bg-gray-50">
            <div className="flex flex-col max-w-full md:max-w-[375px] w-full md:w-[375px] h-full md:h-[812px] max-h-full md:max-h-[812px] border border-gray-200 bg-white">
               <Header />
               <Outlet /> {/* Start rendering router matches */}
               <TanStackRouterDevtools position="bottom-left" />
            </div>
         </main>
      );
   },
});

const homeRoute = new Route({
   getParentRoute: () => rootRoute,
   path: "/",
   component: () => <Navigate to="/payment" />,
});

const paymentRoute = new Route({
   getParentRoute: () => rootRoute,
   path: "payment",
   component: () => <Payment />,
});

const accountRoute = new Route({
   getParentRoute: () => rootRoute,
   path: "payment/account",
   component: () => <Account />,
   validateSearch: AccountRouteSchema,
});

const accountExpiredRoute = new Route({
   getParentRoute: () => rootRoute,
   path: "payment/account/expired",
   component: () => <AccountExpired />,
   validateSearch: AccountRouteSchema,
});

const paymentSuccessRoute = new Route({
   getParentRoute: () => rootRoute,
   path: "payment/success",
   component: () => <PaymentSuccess />,
   validateSearch: AccountRouteSchema,
});

const routeTree = rootRoute.addChildren([
   homeRoute,
   paymentRoute,
   accountRoute,
   accountExpiredRoute,
   paymentSuccessRoute,
]);

const router = new Router({
   routeTree,
   defaultPreload: "intent",
   context: {
      queryClient,
   },
});

declare module "@tanstack/react-router" {
   interface Register {
      router: typeof router;
   }
}

function App() {
   return (
      <QueryClientProvider client={queryClient}>
         <RouterProvider router={router} />
         <ReactQueryDevtools initialIsOpen position="bottom-right" />
      </QueryClientProvider>
   );
}

export default App;

const Header = () => {
   const router = useRouter();

   return (
      <header className="w-full px-5 py-5 bg-primary">
         {router.state.resolvedLocation.pathname == "/payment/account" ? (
            <button
               className="inline-block mb-5"
               onClick={() => router.navigate({ to: ".." })}
            >
               <FontAwesomeIcon
                  icon={faArrowLeftLong}
                  className="w-5 h-5 text-gray-50"
               />
            </button>
         ) : (
            <span className="inline-block mb-5 w-2 h-2" />
         )}
         <div className="flex justify-between items-center">
            <h2 className="text-2xl font-bold text-gray-50">Payment</h2>
            <span>
               <FontAwesomeIcon
                  icon={faBuildingColumns}
                  className="w-5 h-5 text-gray-50"
               />
            </span>
         </div>
      </header>
   );
};

const Payment = () => {
   const [amount, setAmount] = React.useState("");
   const [showInputError, setShowInputError] = React.useState(false);
   const handleGeneratePaymentAccount = mutation.generatePaymentAccount();

   useTimeout(
      () => {
         setShowInputError(false);
      },
      showInputError ? 3000 : null
   );

   return (
      <div className="relative w-full h-full pt-10 pb-5 px-5">
         {showInputError ? (
            <div className="absolute top-0 left-0 w-full px-5 py-2 bg-red-100">
               <p className="text-xs text-red-400 text-center">
                  Please enter a valid amount
               </p>
            </div>
         ) : null}
         <div className="w-full h-full py-2 flex flex-col justify-between">
            <div className="">
               <label
                  htmlFor="amount"
                  className="text-sm text-gray-400 font-medium"
               >
                  Enter amount to pay
               </label>
               <div className="w-full h-9 mt-1 rounded border border-gray-200 bg-gray-50">
                  <input
                     className="appearance-none bg-transparent px-2 w-full h-full text-sm text-gray-600 font-medium"
                     type="number"
                     name="amount"
                     id="amount"
                     placeholder="₦"
                     value={amount}
                     onChange={(e) => setAmount(e.target.value)}
                  />
               </div>
               <div className="flex items-center w-full mt-5 px-5 py-[15px] space-x-2.5 bg-indigo-50 rounded-md">
                  <span className="">
                     <FontAwesomeIcon
                        icon={faInfoCircle}
                        className="w-4 h-4 text-indigo-500"
                     />
                  </span>
                  <p className="text-xs text-indigo-500">
                     A temporary account will be generated for this transaction.
                     You can make a transfer to the account from your choice of
                     bank
                  </p>
               </div>
            </div>
            <div className="w-full self-end">
               <button
                  className="relative flex items-center justify-center w-full h-11 text-gray-50 text-sm bg-primary rounded-md hover:opacity-90"
                  onClick={() => {
                     if (amount === "" || Number.parseFloat(amount) <= 0) {
                        console.log("Please enter a valid amount");
                        setShowInputError(true);
                        return;
                     }
                     handleGeneratePaymentAccount.mutate({
                        amount: Number.parseFloat(amount),
                        session_id: nanoid(12).toUpperCase(),
                     });
                  }}
               >
                  {handleGeneratePaymentAccount.isLoading ? (
                     <LoadingSpinner className="w-10 h-10 text-gray-50" />
                  ) : (
                     <span>Continue</span>
                  )}
                  {!handleGeneratePaymentAccount.isLoading ? (
                     <span className="absolute right-4 z-10">
                        <FontAwesomeIcon
                           icon={faArrowRight}
                           className="w-3.5 h-3.5 text-gray-50"
                        />
                     </span>
                  ) : null}
               </button>
            </div>
         </div>
         <Outlet />
      </div>
   );
};

const Account = () => {
   const [isPolling, setIsPolling] = React.useState(false);
   const { id, txn } = useSearch({ from: accountRoute.id });
   const navigate = useNavigate();
   const fetchPaymentAccountResp = query.fetchPaymentAccount(id);
   const fetchTransactionResp = query.fetchTransaction(txn, isPolling);

   React.useEffect(() => {
      if (fetchTransactionResp.data?.data.status === "SUCCESSFUL") {
         setIsPolling(false);
         navigate({ to: "/payment/success", search: { id, txn } });
      }
   }, [fetchTransactionResp.data?.data.status, id, txn, navigate]);

   return (
      <div className="w-full h-full pt-5 pb-5 flex flex-col">
         {fetchPaymentAccountResp.isLoading ? (
            <React.Fragment>
               <div className="w-full px-5">
                  <p className="text-[14px] text-indigo-500 text-center">
                     We're preparing your account...
                  </p>
               </div>
               <div className="w-full px-5 flex items-center justify-center flex-auto">
                  <LoadingSpinner className="w-14 h-14 text-gray-40" />
               </div>
            </React.Fragment>
         ) : (
            <React.Fragment>
               <div className="w-full flex items-center justify-between bg-indigo-50 px-5 py-2">
                  <p className="text-xs text-indigo-500">
                     <CountdownTimer
                        end={fetchPaymentAccountResp.data?.data?.expires_at}
                        accountId={id}
                        transactionId={txn}
                     />
                  </p>
                  <p className="text-xs text-indigo-500">
                     Pay{" "}
                     <span className="font-medium">
                        NGN{" "}
                        {Intl.NumberFormat("en-US").format(
                           fetchPaymentAccountResp.data?.data.payment_amount ??
                              0
                        )}
                     </span>
                  </p>

                  <span
                     className={`inline-block w-1.5 h-1.5 rounded-full bg-indigo-600 ${
                        fetchTransactionResp.isFetching ? "animate-ping" : ""
                     }`}
                  />
               </div>
               <div className="w-full px-5 flex-auto">
                  <div className="my-6">
                     <p className="text-[15px] text-gray-600">
                        Transfer{" "}
                        <span className="font-semibold">
                           NGN{" "}
                           {Intl.NumberFormat("en-US").format(
                              fetchPaymentAccountResp.data?.data
                                 .payment_amount ?? 0
                           )}
                        </span>{" "}
                        to{" "}
                        <span className="underline">
                           {fetchPaymentAccountResp.data?.data.account_name}
                        </span>
                     </p>
                  </div>
                  <div className="w-full bg-gray-100 border border-gray-200 px-5 py-[15px] rounded-lg space-y-5">
                     <div className="relative space-y-1">
                        <p className="text-xs text-gray-600 uppercase">
                           Bank Name
                        </p>
                        <p className="text-[15px] text-primary font-medium">
                           {fetchPaymentAccountResp.data?.data.bank_name}
                        </p>
                     </div>
                     <div className="relative space-y-1">
                        <p className="text-xs text-gray-600 uppercase">
                           Account Number
                        </p>
                        <p className="text-[15px] text-primary font-medium">
                           {fetchPaymentAccountResp.data?.data.account_number}
                        </p>
                        <span className="absolute right-0 bottom-0 z-10">
                           <CopyToClipboard>
                              {(handleCopy) => (
                                 <button
                                    className="px-0.5"
                                    onClick={() =>
                                       handleCopy(
                                          fetchPaymentAccountResp.data?.data
                                             .account_number ?? ""
                                       )
                                    }
                                 >
                                    <FontAwesomeIcon
                                       icon={faClone}
                                       className="w-3.5 h-3.5 text-gray-400"
                                    />
                                 </button>
                              )}
                           </CopyToClipboard>
                        </span>
                     </div>
                     <div className="relative space-y-1">
                        <p className="text-xs text-gray-600 uppercase">
                           Amount
                        </p>
                        <p className="text-[15px] text-primary font-medium">
                           NGN{" "}
                           {Intl.NumberFormat("en-US").format(
                              fetchPaymentAccountResp.data?.data
                                 .payment_amount ?? 0
                           )}
                        </p>
                        <span className="absolute right-0 bottom-0 z-10">
                           <CopyToClipboard>
                              {(handleCopy) => (
                                 <button
                                    className="px-0.5"
                                    onClick={() =>
                                       handleCopy(
                                          String(
                                             fetchPaymentAccountResp.data?.data
                                                .payment_amount ?? ""
                                          )
                                       )
                                    }
                                 >
                                    <FontAwesomeIcon
                                       icon={faClone}
                                       className="w-3.5 h-3.5 text-gray-400"
                                    />
                                 </button>
                              )}
                           </CopyToClipboard>
                        </span>
                     </div>
                  </div>
                  <div className="flex items-center w-full mt-5 px-5 py-[15px] space-x-2.5 bg-indigo-50 rounded-md">
                     <span className="">
                        <FontAwesomeIcon
                           icon={faInfoCircle}
                           className="w-4 h-4 text-indigo-500"
                        />
                     </span>
                     <p className="text-xs text-indigo-500">
                        You can only use this account for this transaction only.
                        This account is valid for 30 mins
                     </p>
                  </div>
               </div>
               <div className="w-full self-end px-5">
                  <button
                     className="relative flex items-center justify-center w-full h-11 text-gray-50 text-sm bg-primary rounded-md hover:opacity-90"
                     onClick={() => {
                        setIsPolling(true);
                        console.log("I've sent the money");
                     }}
                     disabled={isPolling}
                  >
                     {isPolling ? (
                        <LoadingSpinner className="w-10 h-10 text-gray-50" />
                     ) : (
                        <span> I've sent the money</span>
                     )}
                     {!isPolling ? (
                        <span className="absolute right-4 z-10">
                           <FontAwesomeIcon
                              icon={faArrowRight}
                              className="w-3.5 h-3.5 text-gray-50"
                           />
                        </span>
                     ) : null}
                  </button>
               </div>
            </React.Fragment>
         )}
      </div>
   );
};

const AccountExpired = () => {
   const { id } = useSearch({ from: accountExpiredRoute.id });
   const navigate = useNavigate();
   const fetchPaymentAccountResp = query.fetchPaymentAccount(id);
   const handleGeneratePaymentAccount = mutation.generatePaymentAccount();

   return (
      <div className="w-full h-full pt-5 pb-5 flex flex-col">
         <React.Fragment>
            <div className="w-full flex items-center justify-between bg-amber-50 px-5 py-2">
               <p className="text-xs text-amber-500 line-through">00:00</p>
               <p className="text-xs text-amber-500 line-through">
                  Pay{" "}
                  <span className="font-medium">
                     NGN{" "}
                     {Intl.NumberFormat("en-US").format(
                        fetchPaymentAccountResp.data?.data.payment_amount ?? 0
                     )}
                  </span>
               </p>
               <span className="inline-block">
                  <FontAwesomeIcon
                     icon={faX}
                     className="w-2.5 h-2.5 text-amber-500"
                  />
               </span>
            </div>
            <div className="w-full mt-20 px-5 flex-auto flex flex-col items-center">
               <span className="block">
                  <FontAwesomeIcon
                     icon={faExclamationTriangle}
                     className="w-12 h-12 text-amber-500"
                  />
               </span>
               <div className="mt-[15px]">
                  <p className="text-[15px] text-primary">Account expired</p>
               </div>
               <div className="my-1.5">
                  <p className="text-xs text-gray-600">
                     <span className="underline">
                        {fetchPaymentAccountResp.data?.data.account_name}
                     </span>{" "}
                     is no longer available for this payment
                  </p>
               </div>
            </div>
            <div className="w-full self-end px-5">
               <button
                  className="relative w-full h-11 text-primary text-sm bg-gray-50 border border-gray-200 rounded-md hover:opacity-90"
                  onClick={() => {
                     handleGeneratePaymentAccount.mutate({
                        amount:
                           fetchPaymentAccountResp.data?.data.payment_amount ??
                           0,
                        session_id: nanoid(12).toUpperCase(),
                     });
                  }}
               >
                  {handleGeneratePaymentAccount.isLoading ? (
                     <LoadingSpinner className="w-10 h-10 text-gray-400" />
                  ) : (
                     <span>Try again</span>
                  )}
               </button>
               <button
                  className="relative flex items-center justify-center mt-3 w-full h-11 text-gray-50 text-sm bg-primary rounded-md hover:opacity-90"
                  onClick={() => {
                     navigate({ to: "/" });
                  }}
               >
                  Cancel payment
               </button>
            </div>
         </React.Fragment>
      </div>
   );
};

const PaymentSuccess = () => {
   const { id } = useSearch({ from: paymentSuccessRoute.id });
   const navigate = useNavigate();
   const fetchPaymentAccountResp = query.fetchPaymentAccount(id);

   return (
      <div className="w-full h-full pt-5 pb-5 flex flex-col">
         <React.Fragment>
            <span className="block h-2 w-full" />
            <div className="w-full mt-20 px-5 flex-auto flex flex-col items-center">
               <span className="block">
                  <FontAwesomeIcon
                     icon={faCheckCircle}
                     className="w-12 h-12 text-green-500"
                  />
               </span>
               <div className="mt-[15px]">
                  <p className="text-[15px] text-primary">Payment successful</p>
               </div>
               <div className="my-1.5">
                  <p className="text-xs text-gray-600">
                     <span className="font-medium">
                        NGN{" "}
                        {Intl.NumberFormat("en-US").format(
                           fetchPaymentAccountResp.data?.data.payment_amount ??
                              0
                        )}
                     </span>{" "}
                     has been received by{" "}
                     <span className="underline">
                        {fetchPaymentAccountResp.data?.data.account_name}
                     </span>
                  </p>
               </div>
            </div>
            <div className="w-full self-end px-5">
               <button
                  className="relative w-full h-11 text-gray-50 text-sm bg-primary rounded-md hover:opacity-90"
                  onClick={() => {
                     navigate({ to: "/" });
                  }}
               >
                  Go home
               </button>
            </div>
         </React.Fragment>
      </div>
   );
};

interface CountDownTimerProps {
   end?: Date;
   accountId: string;
   transactionId: string;
}

const CountdownTimer: React.FC<CountDownTimerProps> = ({
   end,
   accountId,
   transactionId,
}) => {
   const [isCounting, setIsCounting] = React.useState(true);
   const [timeLeft, setTimeLeft] = React.useState("");
   const [isRedZone, setIsRedZone] = React.useState(false);
   const navigate = useNavigate();

   useInterval(
      () => {
         const now = dayjs.utc();
         const diff = dayjs(end?.toString().slice(0, 19)).diff(
            now,
            "minutes",
            true
         );

         if (diff < 1) {
            setIsRedZone(true);
         }

         if (diff < 0) {
            setIsCounting(false);
            setTimeLeft("00:00");
            navigate({
               to: "/payment/account/expired",
               search: { id: accountId, txn: transactionId },
            });
            return;
         }

         const minutes = diff - (diff % 1);
         const seconds = Math.round(60 * (diff % 1));
         setTimeLeft(
            `${numberToDoubleDigit(minutes)}:${numberToDoubleDigit(seconds)}`
         );
      },
      isCounting ? 1000 : null
   );

   return (
      <span className={isRedZone ? "text-red-500" : "text-indigo-500"}>
         {timeLeft}
      </span>
   );
};

export const LoadingSpinner: React.FC<{ className: string }> = ({
   className,
}) => (
   <svg
      className={className}
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      viewBox="0 0 100 100"
      preserveAspectRatio="xMidYMid"
   >
      <g transform="rotate(0 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.45s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(36 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.4s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(72 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.35s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(108 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.3s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(144 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.25s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(180 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.2s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(216 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.15s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(252 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.1s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(288 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="-0.05s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
      <g transform="rotate(324 50 50)">
         <rect
            x="48"
            y="26"
            rx="1.68"
            ry="1.68"
            width="4"
            height="12"
            fill="currentColor"
         >
            <animate
               attributeName="opacity"
               values="1;0"
               keyTimes="0;1"
               dur="0.5s"
               begin="0s"
               repeatCount="indefinite"
            ></animate>
         </rect>
      </g>
   </svg>
);

const numberToDoubleDigit = (num: number): string => {
   return num < 10 ? `0${num}` : `${num}`;
};

type CopyFn = (text: string) => Promise<boolean>;

interface CopyToClipboardProps {
   children: (handleCopy: CopyFn) => JSX.Element;
}

const CopyToClipboard: React.FC<CopyToClipboardProps> = ({ children }) => {
   const [hasCopied, setHasCopied] = React.useState(false);
   const [, copyToClipboard] = useCopyToClipboard();

   useTimeout(
      () => {
         setHasCopied(false);
      },
      hasCopied ? 2000 : null
   );

   const handleCopy: CopyFn = async (text: string) => {
      setHasCopied(true);
      return copyToClipboard(text);
   };

   return (
      <React.Fragment>
         {hasCopied ? (
            <span className="text-xs text-gray-600">Copied!</span>
         ) : (
            children(handleCopy)
         )}
      </React.Fragment>
   );
};
