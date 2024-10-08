
/**
 * Samples for Invoices DownloadByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesDownloadByBillingAccount.json
     */
    /**
     * Sample code: InvoicesDownloadByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesDownloadByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().downloadByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "G123456789",
            "12345678", com.azure.core.util.Context.NONE);
    }
}
