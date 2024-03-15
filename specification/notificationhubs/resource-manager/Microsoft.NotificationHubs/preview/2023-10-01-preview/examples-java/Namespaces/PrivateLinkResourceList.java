
/**
 * Samples for PrivateEndpointConnections ListGroupIds.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/PrivateLinkResourceList.json
     */
    /**
     * Sample code: PrivateEndpointConnections_ListGroupIds.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void privateEndpointConnectionsListGroupIds(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.privateEndpointConnections().listGroupIds("5ktrial", "nh-sdk-ns", com.azure.core.util.Context.NONE);
    }
}
