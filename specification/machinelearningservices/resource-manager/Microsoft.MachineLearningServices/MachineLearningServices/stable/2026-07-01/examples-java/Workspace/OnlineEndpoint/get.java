
/**
 * Samples for OnlineEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/OnlineEndpoint/get.json
     */
    /**
     * Sample code: Get Workspace Online Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineEndpoints().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            com.azure.core.util.Context.NONE);
    }
}
