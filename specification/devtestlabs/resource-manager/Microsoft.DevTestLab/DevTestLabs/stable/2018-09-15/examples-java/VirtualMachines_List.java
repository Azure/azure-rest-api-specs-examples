
/**
 * Samples for VirtualMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachines_List.json
     */
    /**
     * Sample code: VirtualMachines_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().list("resourceGroupName", "{labName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
