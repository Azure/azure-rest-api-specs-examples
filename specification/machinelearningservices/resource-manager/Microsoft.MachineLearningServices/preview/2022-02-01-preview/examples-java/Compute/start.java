import com.azure.core.util.Context;

/** Samples for Compute Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/start.json
     */
    /**
     * Sample code: Start ComputeInstance Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void startComputeInstanceCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().start("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
