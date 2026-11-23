
/**
 * Samples for Topics ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicListByNameSpace.json
     */
    /**
     * Sample code: TopicGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().listByNamespace("Default-ServiceBus-WestUS", "sdk-Namespace-1617", null,
            null, com.azure.core.util.Context.NONE);
    }
}
