
/**
 * Samples for BareMetalMachineKeySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachineKeySets_Delete.json
     */
    /**
     * Sample code: Delete bare metal machine key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deleteBareMetalMachineKeySetOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachineKeySets().delete("resourceGroupName", "clusterName", "bareMetalMachineKeySetName", null,
            null, com.azure.core.util.Context.NONE);
    }
}
