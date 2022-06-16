import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateEndpointConnectionsDelete(
        com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .privateEndpointConnections()
            .delete("resRg", "myIoTCentralApp", "myIoTCentralAppEndpoint", Context.NONE);
    }
}
