
/**
 * Samples for VirtualMachineRunCommands Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/runCommandExamples/VirtualMachineRunCommand_Delete.json
     */
    /**
     * Sample code: Delete a run command.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteARunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineRunCommands().delete("myResourceGroup", "myVM", "myRunCommand",
            com.azure.core.util.Context.NONE);
    }
}
