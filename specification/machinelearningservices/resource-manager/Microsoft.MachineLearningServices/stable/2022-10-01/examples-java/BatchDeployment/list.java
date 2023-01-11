/** Samples for BatchDeployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchDeployment/list.json
     */
    /**
     * Sample code: List Batch Deployment.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listBatchDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .batchDeployments()
            .list(
                "test-rg", "my-aml-workspace", "testEndpointName", "string", 1, null, com.azure.core.util.Context.NONE);
    }
}
