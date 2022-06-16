import com.azure.core.util.Context;

/** Samples for VirtualMachineRunCommands GetByVirtualMachine. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/GetRunCommand.json
     */
    /**
     * Sample code: Get a run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .getByVirtualMachineWithResponse("myResourceGroup", "myVM", "myRunCommand", null, Context.NONE);
    }
}
