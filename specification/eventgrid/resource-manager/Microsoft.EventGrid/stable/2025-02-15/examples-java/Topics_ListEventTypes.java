
/**
 * Samples for Topics ListEventTypes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Topics_ListEventTypes.
     * json
     */
    /**
     * Sample code: Topics_ListEventTypes.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsListEventTypes(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().listEventTypes("examplerg", "Microsoft.Storage", "storageAccounts", "ExampleStorageAccount",
            com.azure.core.util.Context.NONE);
    }
}
