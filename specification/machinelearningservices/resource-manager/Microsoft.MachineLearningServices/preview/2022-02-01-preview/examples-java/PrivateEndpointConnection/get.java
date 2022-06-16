import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/PrivateEndpointConnection/get.json
     */
    /**
     * Sample code: WorkspaceGetPrivateEndpointConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void workspaceGetPrivateEndpointConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("rg-1234", "testworkspace", "{privateEndpointConnectionName}", Context.NONE);
    }
}
