
/**
 * Samples for ServerlessEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ServerlessEndpoint/get.json
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
