
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/PrivateEndpointConnection/get.json
     */
    /**
     * Sample code: WorkspaceGetPrivateEndpointConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void workspaceGetPrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.privateEndpointConnections().getWithResponse("rg-1234", "testworkspace",
            "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
