
/**
 * Samples for BareMetalMachineKeySets ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachineKeySets_ListByCluster.json
     */
    /**
     * Sample code: List bare metal machine key sets of the cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listBareMetalMachineKeySetsOfTheCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachineKeySets().listByCluster("resourceGroupName", "clusterName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
