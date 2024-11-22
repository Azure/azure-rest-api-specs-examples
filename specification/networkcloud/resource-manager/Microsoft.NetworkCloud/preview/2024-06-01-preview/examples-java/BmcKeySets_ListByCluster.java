
/**
 * Samples for BmcKeySets ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/
     * BmcKeySets_ListByCluster.json
     */
    /**
     * Sample code: List baseboard management controller key sets of the cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBaseboardManagementControllerKeySetsOfTheCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bmcKeySets().listByCluster("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE);
    }
}
