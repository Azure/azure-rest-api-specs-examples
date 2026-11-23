
/**
 * Samples for Topics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Topics/SBTopicGet.json
     */
    /**
     * Sample code: TopicGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void topicGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getTopics().getWithResponse("ArunMonocle", "sdk-Namespace-1617", "sdk-Topics-5488",
            com.azure.core.util.Context.NONE);
    }
}
