import com.azure.core.util.Context;

/** Samples for BatchDeployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/list.json
     */
    /**
     * Sample code: List Batch Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listBatchDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .batchDeployments()
            .list("test-rg", "my-aml-workspace", "testEndpointName", "string", 1, null, Context.NONE);
    }
}
