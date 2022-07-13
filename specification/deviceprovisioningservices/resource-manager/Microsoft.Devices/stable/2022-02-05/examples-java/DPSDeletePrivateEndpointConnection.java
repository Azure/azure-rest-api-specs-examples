import com.azure.core.util.Context;

/** Samples for IotDpsResource DeletePrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Delete.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void privateEndpointConnectionDelete(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .deletePrivateEndpointConnection(
                "myResourceGroup", "myFirstProvisioningService", "myPrivateEndpointConnection", Context.NONE);
    }
}
