
/**
 * Samples for Transactions GetTransactionSummaryByInvoice.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * transactionSummaryGetByInvoice.json
     */
    /**
     * Sample code: TransactionSummaryGetByInvoice.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void transactionSummaryGetByInvoice(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.transactions().getTransactionSummaryByInvoiceWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "G123456789", null,
            null, com.azure.core.util.Context.NONE);
    }
}
