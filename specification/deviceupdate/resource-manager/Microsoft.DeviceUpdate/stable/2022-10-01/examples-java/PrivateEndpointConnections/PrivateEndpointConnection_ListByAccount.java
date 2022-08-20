import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnections/PrivateEndpointConnection_ListByAccount.json
     */
    /**
     * Sample code: PrivateEndpointConnectionList.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateEndpointConnectionList(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnections().listByAccount("test-rg", "contoso", Context.NONE);
    }
}
