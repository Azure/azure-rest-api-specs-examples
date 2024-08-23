
/**
 * Samples for BatchEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/BatchEndpoint/get.json
     */
    /**
     * Sample code: Get Workspace Batch Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            com.azure.core.util.Context.NONE);
    }
}
