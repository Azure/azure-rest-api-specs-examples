/** Samples for Subscriptions ListLocations. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listLocations.json
     */
    /**
     * Sample code: listLocations.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void listLocations(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptions().listLocations("83aa47df-e3e9-49ff-877b-94304bf3d3ad", com.azure.core.util.Context.NONE);
    }
}
