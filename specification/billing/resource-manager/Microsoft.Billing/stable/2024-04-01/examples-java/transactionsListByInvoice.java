
/**
 * Samples for Transactions ListByInvoice.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionsListByInvoice.
     * json
     */
    /**
     * Sample code: TransactionsListByInvoice.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void transactionsListByInvoice(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.transactions().listByInvoice(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "G123456789", null,
            null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
