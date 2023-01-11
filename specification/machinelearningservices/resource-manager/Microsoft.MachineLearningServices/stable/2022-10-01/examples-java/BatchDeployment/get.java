/** Samples for BatchDeployments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchDeployment/get.json
     */
    /**
     * Sample code: Get Batch Deployment.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getBatchDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .batchDeployments()
            .getWithResponse(
                "test-rg",
                "my-aml-workspace",
                "testEndpointName",
                "testDeploymentName",
                com.azure.core.util.Context.NONE);
    }
}
