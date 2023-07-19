/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListPrivateEndpointConnections.json
     */
    /**
     * Sample code: GetPrivateEndpointConnection.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getPrivateEndpointConnection(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.privateEndpointConnections().listWithResponse("res6977", "sto2527", com.azure.core.util.Context.NONE);
    }
}
