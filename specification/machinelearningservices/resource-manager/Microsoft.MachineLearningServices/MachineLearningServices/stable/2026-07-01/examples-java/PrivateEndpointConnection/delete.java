
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/PrivateEndpointConnection/delete.json
     */
    /**
     * Sample code: WorkspacePutPrivateEndpointConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void workspacePutPrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.privateEndpointConnections().deleteWithResponse("rg-1234", "testworkspace",
            "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
