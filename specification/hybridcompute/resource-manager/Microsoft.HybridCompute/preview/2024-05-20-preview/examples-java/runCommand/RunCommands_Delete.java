
/**
 * Samples for MachineRunCommands Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/
     * runCommand/RunCommands_Delete.json
     */
    /**
     * Sample code: Delete a Machine Run Command.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteAMachineRunCommand(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineRunCommands().delete("myResourceGroup", "myMachine", "myRunCommand",
            com.azure.core.util.Context.NONE);
    }
}
