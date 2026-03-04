
/**
 * Samples for Monitors GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Monitors_Get.json
     */
    /**
     * Sample code: Monitors_Get.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsGet(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitors().getByResourceGroupWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
