
/**
 * Samples for BareMetalMachines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * BareMetalMachines_Delete.json
     */
    /**
     * Sample code: Delete bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().delete("resourceGroupName", "bareMetalMachineName",
            com.azure.core.util.Context.NONE);
    }
}
