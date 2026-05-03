
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: GetPrivateEndpointConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getPrivateEndpointConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.privateEndpointConnections().getWithResponse("res6977", "sto2527", "{privateEndpointConnectionName}",
            com.azure.core.util.Context.NONE);
    }
}
