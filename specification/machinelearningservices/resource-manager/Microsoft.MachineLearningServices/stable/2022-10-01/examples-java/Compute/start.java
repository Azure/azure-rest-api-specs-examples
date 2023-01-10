import com.azure.core.util.Context;

/** Samples for Compute Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/start.json
     */
    /**
     * Sample code: Start ComputeInstance Compute.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void startComputeInstanceCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().start("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
