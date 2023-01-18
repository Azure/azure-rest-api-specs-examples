/** Samples for VirtualMachines Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Restart.json
     */
    /**
     * Sample code: VirtualMachines_Restart.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesRestart(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .restart("resourceGroupName", "{labName}", "{vmName}", com.azure.core.util.Context.NONE);
    }
}
