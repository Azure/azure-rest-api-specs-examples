/** Samples for Subscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listSubscriptions.json
     */
    /**
     * Sample code: listSubscriptions.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void listSubscriptions(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptions().list(com.azure.core.util.Context.NONE);
    }
}
