/** Samples for SubscriptionOperation Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2020-09-01/examples/cancelSubscription.json
     */
    /**
     * Sample code: cancelSubscription.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void cancelSubscription(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager
            .subscriptionOperations()
            .cancelWithResponse("83aa47df-e3e9-49ff-877b-94304bf3d3ad", com.azure.core.util.Context.NONE);
    }
}
