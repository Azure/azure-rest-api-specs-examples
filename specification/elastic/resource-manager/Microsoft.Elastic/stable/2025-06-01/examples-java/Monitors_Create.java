
/**
 * Samples for Monitors Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Monitors_Create.json
     */
    /**
     * Sample code: Monitors_Create.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsCreate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitors().define("myMonitor").withRegion((String) null).withExistingResourceGroup("myResourceGroup")
            .create();
    }
}
