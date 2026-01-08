
/**
 * Samples for BillingAccount GetPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/
     * getBillingAccountPolicy.json
     */
    /**
     * Sample code: GetBillingAccountPolicy.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void getBillingAccountPolicy(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.billingAccounts().getPolicyWithResponse("testBillingAccountId", com.azure.core.util.Context.NONE);
    }
}
