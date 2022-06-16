import com.azure.core.util.Context;

/** Samples for VirtualMachineRunCommands ListByVirtualMachine. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/ListRunCommandsInVM.json
     */
    /**
     * Sample code: List run commands in a Virtual Machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRunCommandsInAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .listByVirtualMachine("myResourceGroup", "myVM", null, Context.NONE);
    }
}
