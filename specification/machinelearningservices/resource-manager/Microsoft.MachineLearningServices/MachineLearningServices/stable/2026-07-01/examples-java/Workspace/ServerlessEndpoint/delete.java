
/**
 * Samples for ServerlessEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ServerlessEndpoint/delete.json
     */
    /**
     * Sample code: Delete Workspace Serverless Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceServerlessEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.serverlessEndpoints().delete("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
