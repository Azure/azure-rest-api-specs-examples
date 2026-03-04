
/**
 * Samples for Organizations Resubscribe.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Organizations_Resubscribe.json
     */
    /**
     * Sample code: Organizations_Resubscribe.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void organizationsResubscribe(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.organizations().resubscribe("myResourceGroup", "myMonitor", null, com.azure.core.util.Context.NONE);
    }
}
