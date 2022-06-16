import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("resRg", "myIoTCentralApp", "myIoTCentralAppEndpoint", Context.NONE);
    }
}
