
/**
 * Samples for BareMetalMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachines_Get.json
     */
    /**
     * Sample code: Get bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().getByResourceGroupWithResponse("resourceGroupName", "bareMetalMachineName",
            com.azure.core.util.Context.NONE);
    }
}
