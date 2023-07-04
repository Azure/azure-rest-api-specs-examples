/** Samples for BmcKeySets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BmcKeySets_Get.json
     */
    /**
     * Sample code: Get baseboard management controller key set of cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getBaseboardManagementControllerKeySetOfCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bmcKeySets()
            .getWithResponse("resourceGroupName", "clusterName", "bmcKeySetName", com.azure.core.util.Context.NONE);
    }
}
