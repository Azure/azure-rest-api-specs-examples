
/**
 * Samples for ClusterManagers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ClusterManagers_Delete.json
     */
    /**
     * Sample code: Delete cluster manager.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteClusterManager(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusterManagers().delete("resourceGroupName", "clusterManagerName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
