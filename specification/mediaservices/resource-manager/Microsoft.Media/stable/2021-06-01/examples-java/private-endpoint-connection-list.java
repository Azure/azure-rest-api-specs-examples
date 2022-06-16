import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/private-endpoint-connection-list.json
     */
    /**
     * Sample code: Get all private endpoint connections.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAllPrivateEndpointConnections(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.privateEndpointConnections().listWithResponse("contoso", "contososports", Context.NONE);
    }
}
