
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_Delete.json
     */
    /**
     * Sample code: Delete cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().delete("resourceGroupName", "clusterName", null, null, com.azure.core.util.Context.NONE);
    }
}
