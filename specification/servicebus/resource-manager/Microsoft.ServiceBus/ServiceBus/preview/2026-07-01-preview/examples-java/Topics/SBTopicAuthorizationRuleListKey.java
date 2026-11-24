
/**
 * Samples for Topics ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicAuthorizationRuleListKey.json
     */
    /**
     * Sample code: TopicAuthorizationRuleListKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicAuthorizationRuleListKey(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().listKeysWithResponse("Default-ServiceBus-WestUS", "sdk-Namespace8408",
            "sdk-Topics2075", "sdk-Authrules5067", com.azure.core.util.Context.NONE);
    }
}
