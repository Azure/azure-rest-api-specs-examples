import com.azure.core.util.Context;

/** Samples for Compute Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/get/ComputeInstance.json
     */
    /**
     * Sample code: Get an ComputeInstance.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getAnComputeInstance(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().getWithResponse("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
