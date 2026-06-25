
/**
 * Samples for BareMetalMachines Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_Restart.json
     */
    /**
     * Sample code: Restart bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void restartBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().restart("resourceGroupName", "bareMetalMachineName",
            com.azure.core.util.Context.NONE);
    }
}
