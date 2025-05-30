
/**
 * Samples for Topics ListSharedAccessKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * Topics_ListSharedAccessKeys.json
     */
    /**
     * Sample code: Topics_ListSharedAccessKeys.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsListSharedAccessKeys(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().listSharedAccessKeysWithResponse("examplerg", "exampletopic2",
            com.azure.core.util.Context.NONE);
    }
}
