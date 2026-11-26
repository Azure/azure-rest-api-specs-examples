
/**
 * Samples for BatchEndpoints ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/BatchEndpoint/listKeys.json
     */
    /**
     * Sample code: ListKeys Workspace Batch Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listKeysWorkspaceBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().listKeysWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            com.azure.core.util.Context.NONE);
    }
}
