
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/Clusters_Get.json
     */
    /**
     * Sample code: Get cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().getByResourceGroupWithResponse("resourceGroupName", "clusterName",
            com.azure.core.util.Context.NONE);
    }
}
