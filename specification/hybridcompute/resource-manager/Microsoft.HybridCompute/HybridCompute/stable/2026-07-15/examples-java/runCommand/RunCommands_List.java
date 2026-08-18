
/**
 * Samples for MachineRunCommands List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/runCommand/RunCommands_List.json
     */
    /**
     * Sample code: GET all Machine Run Commands.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAllMachineRunCommands(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineRunCommands().list("myResourceGroup", "myMachine", null, com.azure.core.util.Context.NONE);
    }
}
