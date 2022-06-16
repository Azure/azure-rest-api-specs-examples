import com.azure.core.util.Context;

/** Samples for Compute Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/get/AKSCompute.json
     */
    /**
     * Sample code: Get a AKS Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getAAKSCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().getWithResponse("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
