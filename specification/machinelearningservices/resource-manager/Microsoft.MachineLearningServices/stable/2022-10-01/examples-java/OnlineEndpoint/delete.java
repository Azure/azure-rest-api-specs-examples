import com.azure.core.util.Context;

/** Samples for OnlineEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/delete.json
     */
    /**
     * Sample code: Delete Online Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineEndpoints().delete("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
