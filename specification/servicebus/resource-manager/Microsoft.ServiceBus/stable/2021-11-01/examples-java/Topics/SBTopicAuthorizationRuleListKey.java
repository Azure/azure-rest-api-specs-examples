
/**
 * Samples for Topics ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicAuthorizationRuleListKey.json
     */
    /**
     * Sample code: TopicAuthorizationRuleListKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicAuthorizationRuleListKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().listKeysWithResponse(
            "Default-ServiceBus-WestUS", "sdk-Namespace8408", "sdk-Topics2075", "sdk-Authrules5067",
            com.azure.core.util.Context.NONE);
    }
}
