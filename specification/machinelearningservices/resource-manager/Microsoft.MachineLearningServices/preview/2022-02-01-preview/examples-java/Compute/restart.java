import com.azure.core.util.Context;

/** Samples for Compute Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/restart.json
     */
    /**
     * Sample code: Restart ComputeInstance Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void restartComputeInstanceCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().restart("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
