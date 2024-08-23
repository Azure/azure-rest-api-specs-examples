
/**
 * Samples for BatchDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/BatchDeployment/get.json
     */
    /**
     * Sample code: Get Workspace Batch Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceBatchDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchDeployments().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            "testDeploymentName", com.azure.core.util.Context.NONE);
    }
}
