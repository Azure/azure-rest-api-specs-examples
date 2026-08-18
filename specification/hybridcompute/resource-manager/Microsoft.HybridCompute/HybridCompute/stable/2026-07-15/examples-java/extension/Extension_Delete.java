
/**
 * Samples for MachineExtensions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/extension/Extension_Delete.json
     */
    /**
     * Sample code: Delete a Machine Extension.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteAMachineExtension(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineExtensions().delete("myResourceGroup", "myMachine", "MMA", com.azure.core.util.Context.NONE);
    }
}
