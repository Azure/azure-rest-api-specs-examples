import com.azure.core.util.Context;

/** Samples for IotDpsResource GetPrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Get.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void privateEndpointConnectionGet(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .getPrivateEndpointConnectionWithResponse(
                "myResourceGroup", "myFirstProvisioningService", "myPrivateEndpointConnection", Context.NONE);
    }
}
