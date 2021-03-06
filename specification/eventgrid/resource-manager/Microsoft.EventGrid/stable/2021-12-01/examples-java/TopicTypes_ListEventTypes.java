import com.azure.core.util.Context;

/** Samples for TopicTypes ListEventTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/TopicTypes_ListEventTypes.json
     */
    /**
     * Sample code: TopicTypes_ListEventTypes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicTypesListEventTypes(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicTypes().listEventTypes("Microsoft.Storage.StorageAccounts", Context.NONE);
    }
}
