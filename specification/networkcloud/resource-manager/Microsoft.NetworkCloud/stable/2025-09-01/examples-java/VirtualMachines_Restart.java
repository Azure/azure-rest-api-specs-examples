
/**
 * Samples for VirtualMachines Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * VirtualMachines_Restart.json
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
