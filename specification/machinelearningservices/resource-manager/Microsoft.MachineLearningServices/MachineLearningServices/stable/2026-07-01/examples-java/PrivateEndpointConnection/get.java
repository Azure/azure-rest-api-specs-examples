
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/PrivateEndpointConnection/get.json
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
