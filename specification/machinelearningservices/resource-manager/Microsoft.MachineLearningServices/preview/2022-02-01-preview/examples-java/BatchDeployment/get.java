import com.azure.core.util.Context;

/** Samples for BatchDeployments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/get.json
     */
    /**
     * Sample code: Get Batch Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getBatchDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .batchDeployments()
            .getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", Context.NONE);
    }
}
