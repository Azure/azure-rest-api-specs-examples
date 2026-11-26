
/**
 * Samples for OnlineDeployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/OnlineDeployment/list.json
     */
    /**
     * Sample code: List Online Deployments.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listOnlineDeployments(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().list("test-rg", "my-aml-workspace", "testEndpointName", "string", 1, null,
            com.azure.core.util.Context.NONE);
    }
}
