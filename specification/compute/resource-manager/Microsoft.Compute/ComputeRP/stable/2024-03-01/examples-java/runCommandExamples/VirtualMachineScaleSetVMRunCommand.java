
import com.azure.resourcemanager.compute.models.RunCommandInput;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSetVMs RunCommand.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/
     * VirtualMachineScaleSetVMRunCommand.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_RunCommand.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().runCommand(
            "myResourceGroup", "myVirtualMachineScaleSet", "0", new RunCommandInput()
                .withCommandId("RunPowerShellScript").withScript(Arrays.asList("Write-Host Hello World!")),
            com.azure.core.util.Context.NONE);
    }
}
