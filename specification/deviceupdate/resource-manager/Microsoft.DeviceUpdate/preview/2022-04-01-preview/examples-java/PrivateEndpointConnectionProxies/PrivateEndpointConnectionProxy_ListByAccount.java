import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnectionProxies ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_ListByAccount.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyList.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateEndpointConnectionProxyList(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnectionProxies().listByAccount("test-rg", "contoso", Context.NONE);
    }
}
