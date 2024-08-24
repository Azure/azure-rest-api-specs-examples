
/**
 * Samples for Queues GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/
     * SBQueueAuthorizationRuleGet.json
     */
    /**
     * Sample code: QueueAuthorizationRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueAuthorizationRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().getAuthorizationRuleWithResponse(
            "ArunMonocle", "sdk-Namespace-7982", "sdk-Queues-2317", "sdk-AuthRules-5800",
            com.azure.core.util.Context.NONE);
    }
}
