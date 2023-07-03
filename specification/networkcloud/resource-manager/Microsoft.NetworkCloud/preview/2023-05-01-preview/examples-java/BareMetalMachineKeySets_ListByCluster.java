/** Samples for BareMetalMachineKeySets ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachineKeySets_ListByCluster.json
     */
    /**
     * Sample code: List bare metal machine key sets of the cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBareMetalMachineKeySetsOfTheCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bareMetalMachineKeySets()
            .listByCluster("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE);
    }
}
