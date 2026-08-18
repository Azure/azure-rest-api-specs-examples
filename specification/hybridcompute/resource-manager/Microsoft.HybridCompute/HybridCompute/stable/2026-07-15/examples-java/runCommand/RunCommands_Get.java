
/**
 * Samples for MachineRunCommands Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/runCommand/RunCommands_Get.json
     */
    /**
     * Sample code: Get a Run Command.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getARunCommand(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineRunCommands().getWithResponse("myResourceGroup", "myMachine", "myRunCommand",
            com.azure.core.util.Context.NONE);
    }
}
