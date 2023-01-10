import com.azure.core.util.Context;

/** Samples for BatchEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/list.json
     */
    /**
     * Sample code: List Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().list("test-rg", "my-aml-workspace", 1, null, Context.NONE);
    }
}
