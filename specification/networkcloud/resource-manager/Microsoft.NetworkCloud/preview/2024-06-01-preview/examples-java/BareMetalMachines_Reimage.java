
/**
 * Samples for BareMetalMachines Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/
     * BareMetalMachines_Reimage.json
     */
    /**
     * Sample code: Reimage bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void reimageBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().reimage("resourceGroupName", "bareMetalMachineName",
            com.azure.core.util.Context.NONE);
    }
}
