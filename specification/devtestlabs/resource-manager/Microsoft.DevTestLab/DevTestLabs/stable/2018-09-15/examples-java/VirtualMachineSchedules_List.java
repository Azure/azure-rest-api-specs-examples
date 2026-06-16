
/**
 * Samples for VirtualMachineSchedules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachineSchedules_List.json
     */
    /**
     * Sample code: VirtualMachineSchedules_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachineSchedulesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachineSchedules().list("resourceGroupName", "{labName}", "{vmName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
