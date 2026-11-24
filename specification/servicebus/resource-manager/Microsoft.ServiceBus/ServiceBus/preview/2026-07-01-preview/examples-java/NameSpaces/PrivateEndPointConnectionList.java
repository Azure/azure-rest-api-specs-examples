
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/PrivateEndPointConnectionList.json
     */
    /**
     * Sample code: NameSpaceCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().list("SDK-ServiceBus-4794", "sdk-Namespace-5828",
            com.azure.core.util.Context.NONE);
    }
}
