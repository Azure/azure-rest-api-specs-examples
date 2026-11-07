
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopePrivateLinkResourceGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkResources().getWithResponse("MyResourceGroup", "MyPrivateLinkScope",
            "azuremonitor", com.azure.core.util.Context.NONE);
    }
}
