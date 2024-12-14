
/**
 * Samples for PrivateEndpointConnectionProxies ListByAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_ListByAccount.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyList.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        privateEndpointConnectionProxyList(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnectionProxies().listByAccount("test-rg", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
