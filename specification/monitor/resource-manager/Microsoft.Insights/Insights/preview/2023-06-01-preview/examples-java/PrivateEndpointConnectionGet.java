
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("MyResourceGroup", "MyPrivateLinkScope",
            "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
