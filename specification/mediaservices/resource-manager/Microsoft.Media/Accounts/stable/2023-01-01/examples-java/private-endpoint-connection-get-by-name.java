/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/private-endpoint-connection-get-by-name.json
     */
    /**
     * Sample code: Get private endpoint connection.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getPrivateEndpointConnection(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("contosorg", "contososports", "connectionName1", com.azure.core.util.Context.NONE);
    }
}
