
/**
 * Samples for Consoles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Consoles_Get.json
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
