
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        deletesAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().delete("MyResourceGroup", "MyPrivateLinkScope",
            "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
