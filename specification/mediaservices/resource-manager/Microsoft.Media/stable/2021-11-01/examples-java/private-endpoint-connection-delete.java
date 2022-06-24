import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-endpoint-connection-delete.json
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
            .deleteWithResponse("contoso", "contososports", "connectionName1", Context.NONE);
    }
}
