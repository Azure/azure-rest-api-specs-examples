
/**
 * Samples for PrivateLinkResources ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopePrivateLinkResourceListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkResources().listByPrivateLinkScopeWithResponse("MyResourceGroup",
            "MyPrivateLinkScope", com.azure.core.util.Context.NONE);
    }
}
