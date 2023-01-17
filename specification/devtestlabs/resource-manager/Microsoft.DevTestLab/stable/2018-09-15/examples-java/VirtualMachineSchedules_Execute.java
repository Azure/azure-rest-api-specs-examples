/** Samples for VirtualMachineSchedules Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Execute.json
     */
    /**
     * Sample code: VirtualMachineSchedules_Execute.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachineSchedulesExecute(
        com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachineSchedules()
            .execute("resourceGroupName", "{labName}", "{vmName}", "LabVmsShutdown", com.azure.core.util.Context.NONE);
    }
}
