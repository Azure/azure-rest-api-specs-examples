
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_deleteprivateendpointconnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Delete.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionDelete(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.privateEndpointConnections().delete("myResourceGroup", "testHub", "myPrivateEndpointConnection",
            com.azure.core.util.Context.NONE);
    }
}
