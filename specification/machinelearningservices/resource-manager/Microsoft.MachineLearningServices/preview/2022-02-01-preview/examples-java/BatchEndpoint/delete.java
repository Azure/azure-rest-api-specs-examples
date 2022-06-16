import com.azure.core.util.Context;

/** Samples for BatchEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/delete.json
     */
    /**
     * Sample code: Delete Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteBatchEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.batchEndpoints().delete("resourceGroup-1234", "testworkspace", "testBatchEndpoint", Context.NONE);
    }
}
