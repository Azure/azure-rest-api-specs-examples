
/**
 * Samples for TopicSpaces ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * TopicSpaces_ListByNamespace.json
     */
    /**
     * Sample code: TopicSpaces_ListByNamespace.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicSpacesListByNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicSpaces().listByNamespace("examplerg", "namespace123", null, null,
            com.azure.core.util.Context.NONE);
    }
}
