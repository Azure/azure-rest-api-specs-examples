
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_ListByResourceGroup.json
     */
    /**
     * Sample code: List clusters for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listClustersForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().listByResourceGroup("resourceGroupName", null, null, com.azure.core.util.Context.NONE);
    }
}
