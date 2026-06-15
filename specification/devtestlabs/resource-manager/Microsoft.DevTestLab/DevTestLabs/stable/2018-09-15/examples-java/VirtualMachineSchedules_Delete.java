
/**
 * Samples for VirtualMachineSchedules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachineSchedules_Delete.json
     */
    /**
     * Sample code: VirtualMachineSchedules_Delete.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachineSchedulesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachineSchedules().deleteWithResponse("resourceGroupName", "{labName}", "{vmName}",
            "LabVmsShutdown", com.azure.core.util.Context.NONE);
    }
}
