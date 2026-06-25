
/**
 * Samples for VirtualMachineRunCommands GetByVirtualMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/runCommandExamples/VirtualMachineRunCommand_Get.json
     */
    /**
     * Sample code: Get a run command.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getARunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineRunCommands().getByVirtualMachineWithResponse("myResourceGroup",
            "myVM", "myRunCommand", null, com.azure.core.util.Context.NONE);
    }
}
