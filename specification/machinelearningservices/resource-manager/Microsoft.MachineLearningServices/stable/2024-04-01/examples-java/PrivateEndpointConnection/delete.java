
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/PrivateEndpointConnection/delete.json
     */
    /**
     * Sample code: WorkspaceDeletePrivateEndpointConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void workspaceDeletePrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.privateEndpointConnections().deleteWithResponse("rg-1234", "testworkspace",
            "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
