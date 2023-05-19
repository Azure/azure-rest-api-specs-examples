/** Samples for TopicTypes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicTypes_List.json
     */
    /**
     * Sample code: TopicTypes_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicTypesList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicTypes().list(com.azure.core.util.Context.NONE);
    }
}
