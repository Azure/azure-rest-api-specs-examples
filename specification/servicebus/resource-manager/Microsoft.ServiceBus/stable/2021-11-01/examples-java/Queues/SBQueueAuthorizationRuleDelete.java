
/**
 * Samples for Queues DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/
     * SBQueueAuthorizationRuleDelete.json
     */
    /**
     * Sample code: QueueAuthorizationRuleDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueAuthorizationRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getQueues().deleteAuthorizationRuleWithResponse(
            "ArunMonocle", "sdk-namespace-7982", "sdk-Queues-2317", "sdk-AuthRules-5800",
            com.azure.core.util.Context.NONE);
    }
}
