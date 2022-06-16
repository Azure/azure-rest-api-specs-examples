import com.azure.core.util.Context;

/** Samples for OnlineEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineEndpoint/get.json
     */
    /**
     * Sample code: Get Online Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.onlineEndpoints().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
