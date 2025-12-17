
/**
 * Samples for VirtualMachines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * VirtualMachines_Delete.json
     */
    /**
     * Sample code: Delete virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().delete("resourceGroupName", "virtualMachineName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
