
/**
 * Samples for VirtualMachines Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Get.
     * json
     */
    /**
     * Sample code: VirtualMachines_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().getWithResponse("resourceGroupName", "{labName}", "{vmName}", null,
            com.azure.core.util.Context.NONE);
    }
}
