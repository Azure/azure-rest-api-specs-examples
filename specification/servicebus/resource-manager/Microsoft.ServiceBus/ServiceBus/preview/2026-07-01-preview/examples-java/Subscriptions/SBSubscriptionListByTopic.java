
/**
 * Samples for Subscriptions ListByTopic.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Subscriptions/SBSubscriptionListByTopic.json
     */
    /**
     * Sample code: SubscriptionListByTopic.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void subscriptionListByTopic(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getSubscriptions().listByTopic("ResourceGroup", "sdk-Namespace-1349", "sdk-Topics-8740",
            null, null, com.azure.core.util.Context.NONE);
    }
}
