
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/PrivateEndPointConnectionList.json
     */
    /**
     * Sample code: PrivateEndPointConnectionList.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void privateEndPointConnectionList(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().list("SDK-EventHub-4794", "sdk-Namespace-5828",
            com.azure.core.util.Context.NONE);
    }
}
