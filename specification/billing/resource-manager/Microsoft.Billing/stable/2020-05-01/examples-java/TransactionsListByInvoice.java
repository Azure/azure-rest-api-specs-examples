/** Samples for Transactions ListByInvoice. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/TransactionsListByInvoice.json
     */
    /**
     * Sample code: TransactionsListByInvoice.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void transactionsListByInvoice(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.transactions().listByInvoice("{billingAccountName}", "{invoiceName}", com.azure.core.util.Context.NONE);
    }
}
