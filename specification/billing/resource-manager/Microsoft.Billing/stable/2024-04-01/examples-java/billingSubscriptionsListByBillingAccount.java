
/**
 * Samples for BillingSubscriptions ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionsListByBillingAccount.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", false, false, null,
            null, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
