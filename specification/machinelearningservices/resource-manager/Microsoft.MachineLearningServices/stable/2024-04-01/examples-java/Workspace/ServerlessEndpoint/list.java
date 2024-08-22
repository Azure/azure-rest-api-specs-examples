
/**
 * Samples for ServerlessEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ServerlessEndpoint/list.json
     */
    /**
     * Sample code: List Workspace Serverless Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceServerlessEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.serverlessEndpoints().list("test-rg", "my-aml-workspace", null, com.azure.core.util.Context.NONE);
    }
}
