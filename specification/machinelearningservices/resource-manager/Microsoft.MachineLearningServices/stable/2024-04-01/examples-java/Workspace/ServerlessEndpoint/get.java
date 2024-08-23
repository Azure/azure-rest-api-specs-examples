
/**
 * Samples for ServerlessEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ServerlessEndpoint/get.json
     */
    /**
     * Sample code: Get Workspace Serverless Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceServerlessEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.serverlessEndpoints().getWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
