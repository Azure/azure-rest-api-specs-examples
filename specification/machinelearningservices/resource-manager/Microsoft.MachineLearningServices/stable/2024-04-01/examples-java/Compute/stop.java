
/**
 * Samples for Compute Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/stop.json
     */
    /**
     * Sample code: Stop ComputeInstance Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        stopComputeInstanceCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().stop("testrg123", "workspaces123", "compute123", com.azure.core.util.Context.NONE);
    }
}
