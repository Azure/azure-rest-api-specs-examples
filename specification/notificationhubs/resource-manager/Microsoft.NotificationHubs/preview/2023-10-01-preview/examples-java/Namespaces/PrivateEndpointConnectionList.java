
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        privateEndpointConnectionsList(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.privateEndpointConnections().list("5ktrial", "nh-sdk-ns", com.azure.core.util.Context.NONE);
    }
}
