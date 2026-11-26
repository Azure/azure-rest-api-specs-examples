
/**
 * Samples for Compute Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Compute/start.json
     */
    /**
     * Sample code: Start ComputeInstance Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        startComputeInstanceCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().start("testrg123", "workspaces123", "compute123", com.azure.core.util.Context.NONE);
    }
}
