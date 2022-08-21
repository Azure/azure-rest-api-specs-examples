import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnectionProxies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyDelete.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateEndpointConnectionProxyDelete(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnectionProxies().delete("test-rg", "contoso", "peexample01", Context.NONE);
    }
}
