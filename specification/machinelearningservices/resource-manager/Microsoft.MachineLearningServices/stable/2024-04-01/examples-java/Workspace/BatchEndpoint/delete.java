
/**
 * Samples for BatchEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/BatchEndpoint/delete.json
     */
    /**
     * Sample code: Delete Workspace Batch Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().delete("resourceGroup-1234", "testworkspace", "testBatchEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
