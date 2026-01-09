
/**
 * Samples for SubscriptionOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/
     * getSubscriptionOperation.json
     */
    /**
     * Sample code: getPendingSubscriptionOperations.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void
        getPendingSubscriptionOperations(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.subscriptionOperations().getWithResponse("e4b8d068-f574-462a-a76f-6fa0afc613c9",
            com.azure.core.util.Context.NONE);
    }
}
