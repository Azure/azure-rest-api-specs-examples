/** Samples for BatchEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/get.json
     */
    /**
     * Sample code: Get Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .batchEndpoints()
            .getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", com.azure.core.util.Context.NONE);
    }
}
