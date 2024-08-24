
/**
 * Samples for Subscriptions ListByTopic.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/
     * SBSubscriptionListByTopic.json
     */
    /**
     * Sample code: SubscriptionListByTopic.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void subscriptionListByTopic(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getSubscriptions().listByTopic("ResourceGroup",
            "sdk-Namespace-1349", "sdk-Topics-8740", null, null, com.azure.core.util.Context.NONE);
    }
}
