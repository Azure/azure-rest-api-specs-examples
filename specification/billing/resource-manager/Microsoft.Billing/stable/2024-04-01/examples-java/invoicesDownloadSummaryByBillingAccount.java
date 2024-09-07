
/**
 * Samples for Invoices DownloadSummaryByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesDownloadSummaryByBillingAccount.json
     */
    /**
     * Sample code: InvoicesDownloadSummaryByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        invoicesDownloadSummaryByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().downloadSummaryByBillingAccount("123456789", "G123456789", com.azure.core.util.Context.NONE);
    }
}
