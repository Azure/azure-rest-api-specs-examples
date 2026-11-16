
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
     * @param manager Entry point to EventHubsManager.
     */
    public static void
        nameSpacePrivateEndPointConnectionGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("SDK-EventHub-4794",
            "sdk-Namespace-5828", "privateEndpointConnectionName", com.azure.core.util.Context.NONE);
    }
}
