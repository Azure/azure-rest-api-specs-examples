/** Samples for OnlineEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/get.json
     */
    /**
     * Sample code: Get Online Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .onlineEndpoints()
            .getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", com.azure.core.util.Context.NONE);
    }
}
