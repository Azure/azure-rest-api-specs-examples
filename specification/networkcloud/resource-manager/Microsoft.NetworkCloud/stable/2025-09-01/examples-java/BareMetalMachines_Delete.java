
/**
 * Samples for BareMetalMachines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachines_Delete.json
     */
    /**
     * Sample code: Delete bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().delete("resourceGroupName", "bareMetalMachineName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
