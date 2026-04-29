
/**
 * Samples for VirtualMachineRunCommands ListByVirtualMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/runCommandExamples/VirtualMachineRunCommand_List.json
     */
    /**
     * Sample code: List run commands in a Virtual Machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listRunCommandsInAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineRunCommands().listByVirtualMachine("myResourceGroup", "myVM", null,
            com.azure.core.util.Context.NONE);
    }
}
