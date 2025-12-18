
/**
 * Samples for BareMetalMachineKeySets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachineKeySets_Get.json
     */
    /**
     * Sample code: Get bare metal machine key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        getBareMetalMachineKeySetOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachineKeySets().getWithResponse("resourceGroupName", "clusterName",
            "bareMetalMachineKeySetName", com.azure.core.util.Context.NONE);
    }
}
