
/**
 * Samples for VirtualMachineRunCommands GetByVirtualMachine.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/runCommandExamples/
     * VirtualMachineRunCommand_Get.json
     */
    /**
     * Sample code: Get a run command.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineRunCommands()
            .getByVirtualMachineWithResponse("myResourceGroup", "myVM", "myRunCommand", null,
                com.azure.core.util.Context.NONE);
    }
}
