import com.azure.core.util.Context;

/** Samples for IotDpsResource ListPrivateEndpointConnections. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSListPrivateEndpointConnections.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void privateEndpointConnectionsList(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .listPrivateEndpointConnectionsWithResponse("myResourceGroup", "myFirstProvisioningService", Context.NONE);
    }
}
