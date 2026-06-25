
/**
 * Samples for VirtualMachines Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VirtualMachines_Start.json
     */
    /**
     * Sample code: Start virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void startVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().start("resourceGroupName", "virtualMachineName", com.azure.core.util.Context.NONE);
    }
}
