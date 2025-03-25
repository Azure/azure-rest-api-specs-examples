
/**
 * Samples for TopicTypes ListEventTypes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/TopicTypes_ListEventTypes
     * .json
     */
    /**
     * Sample code: TopicTypes_ListEventTypes.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicTypesListEventTypes(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicTypes().listEventTypes("Microsoft.Storage.StorageAccounts", com.azure.core.util.Context.NONE);
    }
}
