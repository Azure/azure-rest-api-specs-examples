
/**
 * Samples for SystemTopics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/SystemTopics_Delete.json
     */
    /**
     * Sample code: SystemTopics_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().delete("examplerg", "exampleSystemTopic1", com.azure.core.util.Context.NONE);
    }
}
