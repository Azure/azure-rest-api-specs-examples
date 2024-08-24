
/**
 * Samples for Topics GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicAuthorizationRuleGet.json
     */
    /**
     * Sample code: TopicAuthorizationRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicAuthorizationRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().getAuthorizationRuleWithResponse(
            "ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310",
            com.azure.core.util.Context.NONE);
    }
}
