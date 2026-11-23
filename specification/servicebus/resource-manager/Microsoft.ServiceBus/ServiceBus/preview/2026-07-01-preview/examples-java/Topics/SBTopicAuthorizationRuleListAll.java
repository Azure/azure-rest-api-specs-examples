
/**
 * Samples for Topics ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicAuthorizationRuleListAll.json
     */
    /**
     * Sample code: TopicAuthorizationRuleListAll.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicAuthorizationRuleListAll(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().listAuthorizationRules("ArunMonocle", "sdk-Namespace-6261",
            "sdk-Topics-1984", com.azure.core.util.Context.NONE);
    }
}
