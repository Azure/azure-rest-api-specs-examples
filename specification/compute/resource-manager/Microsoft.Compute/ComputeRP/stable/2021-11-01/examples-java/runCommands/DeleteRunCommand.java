import com.azure.core.util.Context;

/** Samples for VirtualMachineRunCommands Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/DeleteRunCommand.json
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
            .delete("myResourceGroup", "myVM", "myRunCommand", Context.NONE);
    }
}
