
/**
 * Samples for Invoices DownloadByBillingSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesDownloadByBillingSubscription.json
     */
    /**
     * Sample code: InvoicesDownloadByBillingSubscription.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesDownloadByBillingSubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().downloadByBillingSubscription("E123456789", "12345678", com.azure.core.util.Context.NONE);
    }
}
