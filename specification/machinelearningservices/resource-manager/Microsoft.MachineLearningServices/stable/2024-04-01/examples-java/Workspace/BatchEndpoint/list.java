
/**
 * Samples for BatchEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/BatchEndpoint/list.json
     */
    /**
     * Sample code: List Workspace Batch Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().list("test-rg", "my-aml-workspace", 1, null, com.azure.core.util.Context.NONE);
    }
}
