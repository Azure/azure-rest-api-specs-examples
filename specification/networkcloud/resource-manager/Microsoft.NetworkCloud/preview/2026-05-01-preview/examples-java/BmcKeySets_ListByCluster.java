
/**
 * Samples for BmcKeySets ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BmcKeySets_ListByCluster.json
     */
    /**
     * Sample code: List baseboard management controller key sets of the cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBaseboardManagementControllerKeySetsOfTheCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bmcKeySets().listByCluster("resourceGroupName", "clusterName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
