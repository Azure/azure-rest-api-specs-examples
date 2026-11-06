
/**
 * Samples for PrivateEndpointConnections ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a private link scope.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsListOfPrivateEndpointConnectionsOnAPrivateLinkScope(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().listByPrivateLinkScope("MyResourceGroup",
            "MyPrivateLinkScope", com.azure.core.util.Context.NONE);
    }
}
