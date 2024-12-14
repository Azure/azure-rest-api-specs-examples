
/**
 * Samples for PrivateEndpointConnectionProxies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnectionProxyGet.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        privateEndpointConnectionProxyGet(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnectionProxies().getWithResponse("test-rg", "contoso", "peexample01",
            com.azure.core.util.Context.NONE);
    }
}
