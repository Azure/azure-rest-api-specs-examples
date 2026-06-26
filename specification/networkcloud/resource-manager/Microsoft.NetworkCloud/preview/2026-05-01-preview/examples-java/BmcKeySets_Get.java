
/**
 * Samples for BmcKeySets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BmcKeySets_Get.json
     */
    /**
     * Sample code: Get baseboard management controller key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getBaseboardManagementControllerKeySetOfCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bmcKeySets().getWithResponse("resourceGroupName", "clusterName", "bmcKeySetName",
            com.azure.core.util.Context.NONE);
    }
}
