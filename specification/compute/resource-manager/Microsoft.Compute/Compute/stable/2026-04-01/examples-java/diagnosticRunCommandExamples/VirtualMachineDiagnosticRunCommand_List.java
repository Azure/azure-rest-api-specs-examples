
/**
 * Samples for VirtualMachineDiagnosticRunCommands DiagnosticListByVirtualMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/diagnosticRunCommandExamples/VirtualMachineDiagnosticRunCommand_List.json
     */
    /**
     * Sample code: List diagnostic run commands in a Virtual Machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listDiagnosticRunCommandsInAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineDiagnosticRunCommands()
            .diagnosticListByVirtualMachine("myResourceGroup", "myVM", null, com.azure.core.util.Context.NONE);
    }
}
