
/**
 * Samples for Datastores Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Datastores_Get.json
     */
    /**
     * Sample code: Datastores_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void datastoresGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.datastores().getWithResponse("group1", "cloud1", "cluster1", "datastore1",
            com.azure.core.util.Context.NONE);
    }
}
