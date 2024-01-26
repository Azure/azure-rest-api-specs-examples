
/**
 * Samples for VirtualMachineRunCommands List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/runCommandExamples/
     * RunCommand_List.json
     */
    /**
     * Sample code: VirtualMachineRunCommandList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineRunCommandList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineRunCommands().list("SoutheastAsia",
            com.azure.core.util.Context.NONE);
    }
}
