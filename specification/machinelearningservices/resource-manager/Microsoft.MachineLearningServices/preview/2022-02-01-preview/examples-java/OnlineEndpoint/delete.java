import com.azure.core.util.Context;

/** Samples for OnlineEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineEndpoint/delete.json
     */
    /**
     * Sample code: Delete Online Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.onlineEndpoints().delete("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
