
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * PrivateEndpointConnections/PrivateEndpointConnection_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnectionDelete.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        privateEndpointConnectionDelete(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnections().delete("test-rg", "contoso", "peexample01",
            com.azure.core.util.Context.NONE);
    }
}
