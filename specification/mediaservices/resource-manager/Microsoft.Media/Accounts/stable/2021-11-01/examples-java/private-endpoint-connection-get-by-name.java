/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/private-endpoint-connection-get-by-name.json
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
            .getWithResponse("contoso", "contososports", "connectionName1", com.azure.core.util.Context.NONE);
    }
}
