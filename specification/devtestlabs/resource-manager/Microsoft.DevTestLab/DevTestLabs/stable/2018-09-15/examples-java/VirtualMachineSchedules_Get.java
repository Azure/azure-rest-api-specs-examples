
/**
 * Samples for VirtualMachineSchedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachineSchedules_Get.json
     */
    /**
     * Sample code: VirtualMachineSchedules_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachineSchedulesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachineSchedules().getWithResponse("resourceGroupName", "{labName}", "{vmName}",
            "LabVmsShutdown", null, com.azure.core.util.Context.NONE);
    }
}
