
/**
 * Samples for VirtualMachines Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * VirtualMachines_Reimage.json
     */
    /**
     * Sample code: Reimage virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void reimageVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().reimage("resourceGroupName", "virtualMachineName", com.azure.core.util.Context.NONE);
    }
}
