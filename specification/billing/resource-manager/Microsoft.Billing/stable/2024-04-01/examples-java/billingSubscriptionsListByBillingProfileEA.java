
/**
 * Samples for BillingSubscriptions ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionsListByBillingProfileEA.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingProfileEA.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionsListByBillingProfileEA(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().listByBillingProfile("pcn.94077792", "6478903", null, null, null, null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
