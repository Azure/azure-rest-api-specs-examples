
/**
 * Samples for VirtualMachineDiagnosticRunCommands Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/diagnosticRunCommandExamples/VirtualMachineDiagnosticRunCommand_Delete.json
     */
    /**
     * Sample code: Delete a diagnostic run command.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteADiagnosticRunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineDiagnosticRunCommands().delete("myResourceGroup", "myVM",
            "myRunCommand", com.azure.core.util.Context.NONE);
    }
}
