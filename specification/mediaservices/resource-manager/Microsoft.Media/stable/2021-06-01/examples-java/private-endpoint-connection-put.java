import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.PrivateEndpointConnection;
import com.azure.resourcemanager.mediaservices.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.mediaservices.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/private-endpoint-connection-put.json
     */
    /**
     * Sample code: Update private endpoint connection.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updatePrivateEndpointConnection(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        PrivateEndpointConnection resource =
            manager
                .privateEndpointConnections()
                .getWithResponse("contoso", "contososports", "connectionName1", Context.NONE)
                .getValue();
        resource
            .update()
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Test description."))
            .apply();
    }
}
