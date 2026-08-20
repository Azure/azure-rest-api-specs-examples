
/**
 * Samples for VirtualMachineScaleSetVMDiagnosticRunCommands Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/diagnosticRunCommandExamples/VirtualMachineScaleSetVMDiagnosticRunCommand_Delete.json
     */
    /**
     * Sample code: Delete VirtualMachineScaleSet VM diagnostic run command.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        deleteVirtualMachineScaleSetVMDiagnosticRunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMDiagnosticRunCommands().delete("myResourceGroup",
            "myvmScaleSet", "0", "myRunCommand", com.azure.core.util.Context.NONE);
    }
}
