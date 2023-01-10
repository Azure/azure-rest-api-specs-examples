import com.azure.core.util.Context;

/** Samples for BatchEndpoints ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/listKeys.json
     */
    /**
     * Sample code: ListKeys Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listKeysBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().listKeysWithResponse("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
