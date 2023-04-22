/** Samples for VirtualMachineRunCommands Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/runCommandExamples/VirtualMachineRunCommand_Delete.json
     */
    /**
     * Sample code: Delete a run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteARunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .delete("myResourceGroup", "myVM", "myRunCommand", com.azure.core.util.Context.NONE);
    }
}
