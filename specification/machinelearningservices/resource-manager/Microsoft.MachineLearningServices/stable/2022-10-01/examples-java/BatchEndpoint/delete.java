import com.azure.core.util.Context;

/** Samples for BatchEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/delete.json
     */
    /**
     * Sample code: Delete Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteBatchEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.batchEndpoints().delete("resourceGroup-1234", "testworkspace", "testBatchEndpoint", Context.NONE);
    }
}
