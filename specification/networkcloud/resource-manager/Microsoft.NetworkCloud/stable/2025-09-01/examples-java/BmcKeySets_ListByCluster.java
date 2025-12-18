
/**
 * Samples for BmcKeySets ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BmcKeySets_ListByCluster.json
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
