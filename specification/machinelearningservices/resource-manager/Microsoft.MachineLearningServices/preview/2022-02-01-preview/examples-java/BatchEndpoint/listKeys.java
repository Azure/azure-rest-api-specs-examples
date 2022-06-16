import com.azure.core.util.Context;

/** Samples for BatchEndpoints ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/listKeys.json
     */
    /**
     * Sample code: ListKeys Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listKeysBatchEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.batchEndpoints().listKeysWithResponse("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
