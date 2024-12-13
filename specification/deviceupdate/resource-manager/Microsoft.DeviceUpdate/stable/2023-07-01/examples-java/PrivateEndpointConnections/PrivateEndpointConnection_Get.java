
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * PrivateEndpointConnections/PrivateEndpointConnection_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnectionGet.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        privateEndpointConnectionGet(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnections().getWithResponse("test-rg", "contoso", "peexample01",
            com.azure.core.util.Context.NONE);
    }
}
