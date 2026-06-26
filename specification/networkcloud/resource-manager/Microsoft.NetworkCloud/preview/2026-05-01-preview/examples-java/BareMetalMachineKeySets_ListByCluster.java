
/**
 * Samples for BareMetalMachineKeySets ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachineKeySets_ListByCluster.json
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
