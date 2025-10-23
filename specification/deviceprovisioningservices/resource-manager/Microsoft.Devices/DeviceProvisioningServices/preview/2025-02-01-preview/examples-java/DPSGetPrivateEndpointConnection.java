
/**
 * Samples for IotDpsResource GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Get.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        privateEndpointConnectionGet(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().getPrivateEndpointConnectionWithResponse("myResourceGroup",
            "myFirstProvisioningService", "myPrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
