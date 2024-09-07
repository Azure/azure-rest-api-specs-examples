
/**
 * Samples for BillingRequests ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRequestsListByBillingAccount.json
     */
    /**
     * Sample code: BillingRequestsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRequestsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRequests().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", null, null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
