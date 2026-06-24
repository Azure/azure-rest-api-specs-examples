
import com.azure.resourcemanager.compute.models.RunCommandInput;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSetVMs RunCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/runCommandExamples/VirtualMachineScaleSetVMRunCommand.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_RunCommand.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetVMsRunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().runCommand(
            "myResourceGroup", "myVirtualMachineScaleSet", "0", new RunCommandInput()
                .withCommandId("RunPowerShellScript").withScript(Arrays.asList("Write-Host Hello World!")),
            com.azure.core.util.Context.NONE);
    }
}
