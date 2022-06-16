import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: DeletePrivateEndpointConnection.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deletePrivateEndpointConnection(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .privateEndpointConnections()
            .delete("res6977", "sto2527", "{privateEndpointConnectionName}", Context.NONE);
    }
}
