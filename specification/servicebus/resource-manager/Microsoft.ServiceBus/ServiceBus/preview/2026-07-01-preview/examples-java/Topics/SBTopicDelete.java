
/**
 * Samples for Topics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicDelete.json
     */
    /**
     * Sample code: TopicDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().deleteWithResponse("ArunMonocle", "sdk-Namespace-1617", "sdk-Topics-5488",
            com.azure.core.util.Context.NONE);
    }
}
