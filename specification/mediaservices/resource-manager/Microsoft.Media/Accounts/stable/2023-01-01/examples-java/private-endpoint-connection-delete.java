/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/private-endpoint-connection-delete.json
     */
    /**
     * Sample code: Delete private endpoint connection.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deletePrivateEndpointConnection(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .privateEndpointConnections()
            .deleteWithResponse("contosorg", "contososports", "connectionName1", com.azure.core.util.Context.NONE);
    }
}
