
/**
 * Samples for Topics DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicAuthorizationRuleDelete.json
     */
    /**
     * Sample code: TopicAuthorizationRuleDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicAuthorizationRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().deleteAuthorizationRuleWithResponse(
            "ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310",
            com.azure.core.util.Context.NONE);
    }
}
