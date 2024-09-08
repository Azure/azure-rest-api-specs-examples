
/**
 * Samples for Policies GetBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesGetBySubscription.
     * json
     */
    /**
     * Sample code: PoliciesGetBySubscription.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void policiesGetBySubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.policies().getBySubscriptionWithResponse(com.azure.core.util.Context.NONE);
    }
}
