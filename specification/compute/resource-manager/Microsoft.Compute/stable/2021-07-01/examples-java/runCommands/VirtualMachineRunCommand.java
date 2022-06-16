import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.RunCommandInput;

/** Samples for VirtualMachines RunCommand. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/VirtualMachineRunCommand.json
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
                "crptestar98131", "vm3036", new RunCommandInput().withCommandId("RunPowerShellScript"), Context.NONE);
    }
}
