
/**
 * Samples for BareMetalMachineKeySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/
     * BareMetalMachineKeySets_Delete.json
     */
    /**
     * Sample code: Delete bare metal machine key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deleteBareMetalMachineKeySetOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachineKeySets().delete("resourceGroupName", "clusterName", "bareMetalMachineKeySetName",
            com.azure.core.util.Context.NONE);
    }
}
