import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/PrivateEndpointConnection/delete.json
     */
    /**
     * Sample code: WorkspaceDeletePrivateEndpointConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void workspaceDeletePrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .privateEndpointConnections()
            .deleteWithResponse("rg-1234", "testworkspace", "{privateEndpointConnectionName}", Context.NONE);
    }
}
