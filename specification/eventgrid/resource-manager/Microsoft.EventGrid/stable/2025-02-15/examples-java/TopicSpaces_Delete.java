
/**
 * Samples for TopicSpaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/TopicSpaces_Delete.json
     */
    /**
     * Sample code: TopicSpaces_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicSpacesDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicSpaces().delete("examplerg", "exampleNamespaceName1", "exampleTopicSpaceName1",
            com.azure.core.util.Context.NONE);
    }
}
