/** Samples for VirtualMachineRunCommands Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/runCommandExamples/RunCommand_Get.json
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
            .getWithResponse("SoutheastAsia", "RunPowerShellScript", com.azure.core.util.Context.NONE);
    }
}
