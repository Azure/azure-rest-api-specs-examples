/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/private-endpoint-connection-list.json
     */
    /**
     * Sample code: Get all private endpoint connections.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAllPrivateEndpointConnections(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .privateEndpointConnections()
            .listWithResponse("contosorg", "contososports", com.azure.core.util.Context.NONE);
    }
}
