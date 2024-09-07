
/**
 * Samples for Invoices GetByBillingSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesGetByBillingSubscription.json
     */
    /**
     * Sample code: InvoicesGetByBillingSubscription.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesGetByBillingSubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getByBillingSubscriptionWithResponse("E123456789", com.azure.core.util.Context.NONE);
    }
}
