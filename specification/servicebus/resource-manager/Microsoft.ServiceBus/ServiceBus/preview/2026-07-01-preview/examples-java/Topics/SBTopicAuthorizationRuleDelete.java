
/**
 * Samples for Topics DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicAuthorizationRuleDelete.json
     */
    /**
     * Sample code: TopicAuthorizationRuleDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicAuthorizationRuleDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().deleteAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-6261",
            "sdk-Topics-1984", "sdk-AuthRules-4310", com.azure.core.util.Context.NONE);
    }
}
