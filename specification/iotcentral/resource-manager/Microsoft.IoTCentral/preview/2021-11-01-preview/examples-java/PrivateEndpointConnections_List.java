import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_List.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateEndpointConnectionsList(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.privateEndpointConnections().list("resRg", "myIoTCentralApp", Context.NONE);
    }
}
