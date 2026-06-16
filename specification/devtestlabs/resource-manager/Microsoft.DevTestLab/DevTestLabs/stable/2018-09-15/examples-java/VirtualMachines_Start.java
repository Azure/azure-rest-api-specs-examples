
/**
 * Samples for VirtualMachines Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachines_Start.json
     */
    /**
     * Sample code: VirtualMachines_Start.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesStart(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().start("resourceGroupName", "{labName}", "{vmName}", com.azure.core.util.Context.NONE);
    }
}
