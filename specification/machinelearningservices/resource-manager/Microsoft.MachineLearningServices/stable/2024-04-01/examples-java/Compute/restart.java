
/**
 * Samples for Compute Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/restart.json
     */
    /**
     * Sample code: Restart ComputeInstance Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        restartComputeInstanceCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().restart("testrg123", "workspaces123", "compute123", com.azure.core.util.Context.NONE);
    }
}
