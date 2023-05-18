/** Samples for ClusterManagers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/ClusterManagers_Get.json
     */
    /**
     * Sample code: Get cluster manager.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getClusterManager(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .clusterManagers()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "clusterManagerName", com.azure.core.util.Context.NONE);
    }
}
