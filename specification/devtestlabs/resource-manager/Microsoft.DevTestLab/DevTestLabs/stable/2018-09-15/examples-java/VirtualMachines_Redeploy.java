
/**
 * Samples for VirtualMachines Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachines_Redeploy.json
     */
    /**
     * Sample code: VirtualMachines_Redeploy.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesRedeploy(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().redeploy("resourceGroupName", "{labName}", "{vmName}",
            com.azure.core.util.Context.NONE);
    }
}
