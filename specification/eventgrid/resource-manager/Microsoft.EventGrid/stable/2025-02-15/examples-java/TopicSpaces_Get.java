
/**
 * Samples for TopicSpaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/TopicSpaces_Get.json
     */
    /**
     * Sample code: TopicSpaces_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicSpacesGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicSpaces().getWithResponse("examplerg", "exampleNamespaceName1", "exampleTopicSpaceName1",
            com.azure.core.util.Context.NONE);
    }
}
