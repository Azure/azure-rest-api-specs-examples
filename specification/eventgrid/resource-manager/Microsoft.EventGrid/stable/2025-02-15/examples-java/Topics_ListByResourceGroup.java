
/**
 * Samples for Topics ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * Topics_ListByResourceGroup.json
     */
    /**
     * Sample code: Topics_ListByResourceGroup.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().listByResourceGroup("examplerg", null, null, com.azure.core.util.Context.NONE);
    }
}
