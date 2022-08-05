import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandScriptSource;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandUpdate;

/** Samples for VirtualMachineRunCommands Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/UpdateRunCommand.json
     */
    /**
     * Sample code: Update a run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateARunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineRunCommands()
            .update(
                "myResourceGroup",
                "myVM",
                "myRunCommand",
                new VirtualMachineRunCommandUpdate()
                    .withSource(
                        new VirtualMachineRunCommandScriptSource().withScript("Write-Host Script Source Updated!")),
                Context.NONE);
    }
}
