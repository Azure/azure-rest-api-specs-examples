/** Samples for VirtualMachineSchedules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_List.json
     */
    /**
     * Sample code: VirtualMachineSchedules_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachineSchedulesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachineSchedules()
            .list(
                "resourceGroupName", "{labName}", "{vmName}", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
