import com.azure.core.util.Context;

/** Samples for VirtualMachineRunCommands Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/VirtualMachineRunCommandGet.json
     */
    /**
     * Sample code: VirtualMachineRunCommandGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineRunCommandGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .getWithResponse("SoutheastAsia", "RunPowerShellScript", Context.NONE);
    }
}
