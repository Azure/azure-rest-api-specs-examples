import com.azure.core.util.Context;

/** Samples for VirtualMachineRunCommands List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/VirtualMachineRunCommandList.json
     */
    /**
     * Sample code: VirtualMachineRunCommandList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineRunCommandList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .list("SoutheastAsia", Context.NONE);
    }
}
