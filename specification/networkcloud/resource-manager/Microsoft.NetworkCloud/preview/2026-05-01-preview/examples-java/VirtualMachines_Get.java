
/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VirtualMachines_Get.json
     */
    /**
     * Sample code: Get virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().getByResourceGroupWithResponse("resourceGroupName", "virtualMachineName",
            com.azure.core.util.Context.NONE);
    }
}
