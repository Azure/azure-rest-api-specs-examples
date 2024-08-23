
/**
 * Samples for Compute Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/get/AKSCompute.json
     */
    /**
     * Sample code: Get a AKS Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getAAKSCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().getWithResponse("testrg123", "workspaces123", "compute123",
            com.azure.core.util.Context.NONE);
    }
}
