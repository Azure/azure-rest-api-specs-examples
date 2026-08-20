
/**
 * Samples for VirtualMachineScaleSetVMDiagnosticRunCommands DiagnosticList.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/diagnosticRunCommandExamples/VirtualMachineScaleSetVMDiagnosticRunCommand_List.json
     */
    /**
     * Sample code: List diagnostic run commands in Vmss instance.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listDiagnosticRunCommandsInVmssInstance(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMDiagnosticRunCommands().diagnosticList("myResourceGroup",
            "myvmScaleSet", "0", null, com.azure.core.util.Context.NONE);
    }
}
