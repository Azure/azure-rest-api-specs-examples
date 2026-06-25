
/**
 * Samples for VirtualMachines Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VirtualMachines_Restart.json
     */
    /**
     * Sample code: Restart virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void restartVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().restart("resourceGroupName", "virtualMachineName", com.azure.core.util.Context.NONE);
    }
}
