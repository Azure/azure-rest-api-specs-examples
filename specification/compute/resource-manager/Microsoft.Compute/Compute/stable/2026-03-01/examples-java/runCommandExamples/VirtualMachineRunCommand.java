
import com.azure.resourcemanager.compute.models.RunCommandInput;

/**
 * Samples for VirtualMachines RunCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/runCommandExamples/VirtualMachineRunCommand.json
     */
    /**
     * Sample code: VirtualMachineRunCommand.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().runCommand("crptestar98131", "vm3036",
            new RunCommandInput().withCommandId("RunPowerShellScript"), com.azure.core.util.Context.NONE);
    }
}
