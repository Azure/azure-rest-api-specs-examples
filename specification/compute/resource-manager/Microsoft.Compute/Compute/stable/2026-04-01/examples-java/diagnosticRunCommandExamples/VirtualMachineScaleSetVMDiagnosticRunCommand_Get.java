
/**
 * Samples for VirtualMachineScaleSetVMDiagnosticRunCommands Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/diagnosticRunCommandExamples/VirtualMachineScaleSetVMDiagnosticRunCommand_Get.json
     */
    /**
     * Sample code: Get VirtualMachineScaleSet VM diagnostic run commands.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getVirtualMachineScaleSetVMDiagnosticRunCommands(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMDiagnosticRunCommands().getWithResponse("myResourceGroup",
            "myvmScaleSet", "0", "myRunCommand", null, com.azure.core.util.Context.NONE);
    }
}
