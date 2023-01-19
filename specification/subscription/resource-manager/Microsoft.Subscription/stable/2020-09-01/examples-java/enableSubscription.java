/** Samples for SubscriptionOperation Enable. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2020-09-01/examples/enableSubscription.json
     */
    /**
     * Sample code: enableSubscription.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void enableSubscription(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager
            .subscriptionOperations()
            .enableWithResponse("7948bcee-488c-47ce-941c-38e20ede803d", com.azure.core.util.Context.NONE);
    }
}
