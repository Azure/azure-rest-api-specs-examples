/** Samples for Subscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/getSubscription.json
     */
    /**
     * Sample code: getSubscription.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void getSubscription(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager
            .subscriptions()
            .getWithResponse("83aa47df-e3e9-49ff-877b-94304bf3d3ad", com.azure.core.util.Context.NONE);
    }
}
