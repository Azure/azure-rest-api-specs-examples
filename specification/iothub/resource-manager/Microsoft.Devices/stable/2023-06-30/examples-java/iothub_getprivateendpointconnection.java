
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/
     * iothub_getprivateendpointconnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Get.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionGet(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.privateEndpointConnections().getWithResponse("myResourceGroup", "testHub",
            "myPrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
