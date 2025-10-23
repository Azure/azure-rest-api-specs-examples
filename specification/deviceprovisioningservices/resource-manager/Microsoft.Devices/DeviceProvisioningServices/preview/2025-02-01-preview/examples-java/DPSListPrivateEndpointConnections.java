
/**
 * Samples for IotDpsResource ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSListPrivateEndpointConnections.json
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
