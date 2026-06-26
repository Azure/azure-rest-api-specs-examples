
/**
 * Samples for Consoles ListByVirtualMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Consoles_ListByVirtualMachine.json
     */
    /**
     * Sample code: List consoles of the virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listConsolesOfTheVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.consoles().listByVirtualMachine("resourceGroupName", "virtualMachineName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
