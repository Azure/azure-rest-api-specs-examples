
/**
 * Samples for BillingRequests ListByUser.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsListByUser.
     * json
     */
    /**
     * Sample code: BillingRequestsListByUser.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRequestsListByUser(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRequests().listByUser(null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
