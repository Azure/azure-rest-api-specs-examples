
/**
 * Samples for BillingRequests ListByUser.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRequestsListByUserWithFilter.json
     */
    /**
     * Sample code: BillingRequestsListByUserWithFilter.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRequestsListByUserWithFilter(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRequests().listByUser("properties/status eq 'Approved'", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
