
/**
 * Samples for SubscriptionPolicy ListPolicyForTenant.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getTenantPolicyList
     * .json
     */
    /**
     * Sample code: getTenantPolicyList.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void getTenantPolicyList(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptionPolicies().listPolicyForTenant(com.azure.core.util.Context.NONE);
    }
}
