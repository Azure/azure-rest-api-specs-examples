
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListPrivateEndpointConnections.json
     */
    /**
     * Sample code: GetPrivateEndpointConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getPrivateEndpointConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.privateEndpointConnections().listWithResponse("res6977", "sto2527", com.azure.core.util.Context.NONE);
    }
}
