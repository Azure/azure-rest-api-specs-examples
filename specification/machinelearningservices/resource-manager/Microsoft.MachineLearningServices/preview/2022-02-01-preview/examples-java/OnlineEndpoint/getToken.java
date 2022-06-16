import com.azure.core.util.Context;

/** Samples for OnlineEndpoints GetToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineEndpoint/getToken.json
     */
    /**
     * Sample code: GetToken Online Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getTokenOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.onlineEndpoints().getTokenWithResponse("test-rg", "my-aml-workspace", "testEndpointName", Context.NONE);
    }
}
