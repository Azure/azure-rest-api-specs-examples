import com.azure.resourcemanager.compute.models.RunCommandInput;

/** Samples for VirtualMachines RunCommand. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/runCommandExamples/VirtualMachineRunCommand.json
     */
    /**
     * Sample code: VirtualMachineRunCommand.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .runCommand(
                "crptestar98131",
                "vm3036",
                new RunCommandInput().withCommandId("RunPowerShellScript"),
                com.azure.core.util.Context.NONE);
    }
}
