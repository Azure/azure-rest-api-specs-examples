
/**
 * Samples for BillingAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountsList.json
     */
    /**
     * Sample code: BillingAccountsList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().list(null, null, null, null, null, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
