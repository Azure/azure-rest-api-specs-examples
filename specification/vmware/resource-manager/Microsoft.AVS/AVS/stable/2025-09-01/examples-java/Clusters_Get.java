
/**
 * Samples for Clusters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Clusters_Get.json
     */
    /**
     * Sample code: Clusters_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void clustersGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().getWithResponse("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
