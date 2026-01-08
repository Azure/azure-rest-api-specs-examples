
/**
 * Samples for SubscriptionPolicy GetPolicyForTenant.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getTenantPolicy.
     * json
     */
    /**
     * Sample code: getTenantPolicy.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void getTenantPolicy(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptionPolicies().getPolicyForTenantWithResponse(com.azure.core.util.Context.NONE);
    }
}
