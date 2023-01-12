/** Samples for OnlineEndpoints ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/listKeys.json
     */
    /**
     * Sample code: ListKeys Online Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listKeysOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .onlineEndpoints()
            .listKeysWithResponse("test-rg", "my-aml-workspace", "testEndpointName", com.azure.core.util.Context.NONE);
    }
}
