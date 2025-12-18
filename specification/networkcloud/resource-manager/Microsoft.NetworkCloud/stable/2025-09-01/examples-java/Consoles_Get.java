
/**
 * Samples for Consoles Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/Consoles_Get.json
     */
    /**
     * Sample code: Get virtual machine console.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getVirtualMachineConsole(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.consoles().getWithResponse("resourceGroupName", "virtualMachineName", "default",
            com.azure.core.util.Context.NONE);
    }
}
