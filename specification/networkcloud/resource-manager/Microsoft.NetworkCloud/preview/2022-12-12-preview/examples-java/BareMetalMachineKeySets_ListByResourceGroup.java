/** Samples for BareMetalMachineKeySets ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_ListByResourceGroup.json
     */
    /**
     * Sample code: List bare metal machine key set of cluster for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBareMetalMachineKeySetOfClusterForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bareMetalMachineKeySets()
            .listByResourceGroup("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE);
    }
}
