
import com.azure.resourcemanager.notificationhubs.fluent.models.PrivateEndpointConnectionResourceInner;
import com.azure.resourcemanager.notificationhubs.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.notificationhubs.models.PrivateLinkConnectionStatus;
import com.azure.resourcemanager.notificationhubs.models.RemotePrivateEndpointConnection;
import com.azure.resourcemanager.notificationhubs.models.RemotePrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Update.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        privateEndpointConnectionsUpdate(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.privateEndpointConnections().update("5ktrial", "nh-sdk-ns",
            "nh-sdk-ns.1fa229cd-bf3f-47f0-8c49-afb36723997e",
            new PrivateEndpointConnectionResourceInner().withProperties(new PrivateEndpointConnectionProperties()
                .withPrivateEndpoint(new RemotePrivateEndpointConnection()).withPrivateLinkServiceConnectionState(
                    new RemotePrivateLinkServiceConnectionState().withStatus(PrivateLinkConnectionStatus.APPROVED))),
            com.azure.core.util.Context.NONE);
    }
}
