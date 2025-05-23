
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        privateEndpointConnectionsDelete(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.privateEndpointConnections().delete("5ktrial", "nh-sdk-ns",
            "nh-sdk-ns.1fa229cd-bf3f-47f0-8c49-afb36723997e", com.azure.core.util.Context.NONE);
    }
}
