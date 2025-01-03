
/**
 * Samples for IotDpsResource ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/
     * DPSListPrivateEndpointConnections.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        privateEndpointConnectionsList(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listPrivateEndpointConnectionsWithResponse("myResourceGroup",
            "myFirstProvisioningService", com.azure.core.util.Context.NONE);
    }
}
