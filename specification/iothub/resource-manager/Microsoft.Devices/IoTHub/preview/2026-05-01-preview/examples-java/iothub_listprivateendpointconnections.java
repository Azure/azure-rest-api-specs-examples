
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_listprivateendpointconnections.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionsList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.privateEndpointConnections().listWithResponse("myResourceGroup", "testHub",
            com.azure.core.util.Context.NONE);
    }
}
