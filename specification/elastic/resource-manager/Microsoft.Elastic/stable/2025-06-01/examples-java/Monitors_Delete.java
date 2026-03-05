
/**
 * Samples for Monitors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Monitors_Delete.json
     */
    /**
     * Sample code: Monitors_Delete.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsDelete(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitors().delete("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
