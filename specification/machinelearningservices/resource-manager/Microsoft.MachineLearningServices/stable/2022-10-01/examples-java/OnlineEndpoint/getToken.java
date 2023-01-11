/** Samples for OnlineEndpoints GetToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/getToken.json
     */
    /**
     * Sample code: GetToken Online Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getTokenOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .onlineEndpoints()
            .getTokenWithResponse("test-rg", "my-aml-workspace", "testEndpointName", com.azure.core.util.Context.NONE);
    }
}
