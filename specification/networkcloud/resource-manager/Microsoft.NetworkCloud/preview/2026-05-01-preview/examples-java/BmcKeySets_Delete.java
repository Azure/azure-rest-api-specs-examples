
/**
 * Samples for BmcKeySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BmcKeySets_Delete.json
     */
    /**
     * Sample code: Delete baseboard management controller key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteBaseboardManagementControllerKeySetOfCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bmcKeySets().delete("resourceGroupName", "clusterName", "bmcKeySetName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
