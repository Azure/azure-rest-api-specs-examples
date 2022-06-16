import com.azure.core.util.Context;

/** Samples for Compute Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/get/ComputeInstance.json
     */
    /**
     * Sample code: Get an ComputeInstance.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getAnComputeInstance(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().getWithResponse("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
