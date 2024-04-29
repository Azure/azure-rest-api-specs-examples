
/**
 * Samples for SystemTopics List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * SystemTopics_ListBySubscription.json
     */
    /**
     * Sample code: SystemTopics_ListBySubscription.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().list(null, null, com.azure.core.util.Context.NONE);
    }
}
