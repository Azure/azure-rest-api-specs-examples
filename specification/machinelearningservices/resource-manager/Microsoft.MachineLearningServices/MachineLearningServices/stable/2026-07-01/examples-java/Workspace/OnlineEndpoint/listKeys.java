
/**
 * Samples for OnlineEndpoints ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/OnlineEndpoint/listKeys.json
     */
    /**
     * Sample code: ListKeys Workspace Online Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listKeysWorkspaceOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineEndpoints().listKeysWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            com.azure.core.util.Context.NONE);
    }
}
