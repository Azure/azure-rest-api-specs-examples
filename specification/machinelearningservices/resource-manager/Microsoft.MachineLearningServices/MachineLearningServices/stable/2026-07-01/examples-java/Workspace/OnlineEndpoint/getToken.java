
/**
 * Samples for OnlineEndpoints GetToken.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/OnlineEndpoint/getToken.json
     */
    /**
     * Sample code: GetToken Workspace Online Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getTokenWorkspaceOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineEndpoints().getTokenWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            com.azure.core.util.Context.NONE);
    }
}
