import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.RunCommandInput;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSetVMs RunCommand. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/VMScaleSetRunCommand.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_RunCommand.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .runCommand(
                "myResourceGroup",
                "myVirtualMachineScaleSet",
                "0",
                new RunCommandInput()
                    .withCommandId("RunPowerShellScript")
                    .withScript(Arrays.asList("Write-Host Hello World!")),
                Context.NONE);
    }
}
