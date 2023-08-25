/** Samples for ClusterManagers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/ClusterManagers_Delete.json
     */
    /**
     * Sample code: Delete cluster manager.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteClusterManager(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusterManagers().delete("resourceGroupName", "clusterManagerName", com.azure.core.util.Context.NONE);
    }
}
