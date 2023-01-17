/** Samples for VirtualMachines Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Stop.json
     */
    /**
     * Sample code: VirtualMachines_Stop.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesStop(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().stop("resourceGroupName", "{labName}", "{vmName}", com.azure.core.util.Context.NONE);
    }
}
