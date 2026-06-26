
/**
 * Samples for BareMetalMachines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_Delete.json
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
