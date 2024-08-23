
/**
 * Samples for BatchDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/BatchDeployment/delete.json
     */
    /**
     * Sample code: Delete Workspace Batch Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceBatchDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchDeployments().delete("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName",
            com.azure.core.util.Context.NONE);
    }
}
