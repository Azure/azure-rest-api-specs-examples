
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/Clusters_Delete.
     * json
     */
    /**
     * Sample code: Delete cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().delete("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE);
    }
}
