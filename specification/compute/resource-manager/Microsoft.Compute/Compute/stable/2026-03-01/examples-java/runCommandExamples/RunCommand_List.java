
/**
 * Samples for VirtualMachineRunCommands List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/runCommandExamples/RunCommand_List.json
     */
    /**
     * Sample code: VirtualMachineRunCommandList.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRunCommandList(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineRunCommands().list("SoutheastAsia", com.azure.core.util.Context.NONE);
    }
}
