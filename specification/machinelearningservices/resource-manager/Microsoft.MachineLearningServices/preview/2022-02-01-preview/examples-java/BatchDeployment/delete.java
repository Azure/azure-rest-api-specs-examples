import com.azure.core.util.Context;

/** Samples for BatchDeployments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/delete.json
     */
    /**
     * Sample code: Delete Batch Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteBatchDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .batchDeployments()
            .delete("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", Context.NONE);
    }
}
