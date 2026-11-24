
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/PrivateEndPointConnectionGet.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpacePrivateEndPointConnectionGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("SDK-ServiceBus-4794",
            "sdk-Namespace-5828", "privateEndpointConnectionName", com.azure.core.util.Context.NONE);
    }
}
