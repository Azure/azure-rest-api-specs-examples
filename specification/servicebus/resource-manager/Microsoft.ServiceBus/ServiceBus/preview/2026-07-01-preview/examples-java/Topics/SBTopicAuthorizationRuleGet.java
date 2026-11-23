
/**
 * Samples for Topics GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicAuthorizationRuleGet.json
     */
    /**
     * Sample code: TopicAuthorizationRuleGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicAuthorizationRuleGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().getAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-6261",
            "sdk-Topics-1984", "sdk-AuthRules-4310", com.azure.core.util.Context.NONE);
    }
}
