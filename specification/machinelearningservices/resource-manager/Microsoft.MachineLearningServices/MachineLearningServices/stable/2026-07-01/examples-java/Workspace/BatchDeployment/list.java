
/**
 * Samples for BatchDeployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/BatchDeployment/list.json
     */
    /**
     * Sample code: List Workspace Batch Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceBatchDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchDeployments().list("test-rg", "my-aml-workspace", "testEndpointName", "string", 1, null,
            com.azure.core.util.Context.NONE);
    }
}
