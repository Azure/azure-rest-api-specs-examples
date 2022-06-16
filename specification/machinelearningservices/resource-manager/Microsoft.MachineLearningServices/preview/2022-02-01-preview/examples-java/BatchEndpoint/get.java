import com.azure.core.util.Context;

/** Samples for BatchEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/get.json
     */
    /**
     * Sample code: Get Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getBatchEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.batchEndpoints().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
