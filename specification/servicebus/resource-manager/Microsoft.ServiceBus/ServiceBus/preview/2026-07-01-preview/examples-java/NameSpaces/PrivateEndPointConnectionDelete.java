
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/PrivateEndPointConnectionDelete.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpacePrivateEndPointConnectionDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().delete("ArunMonocle", "sdk-Namespace-3285",
            "928c44d5-b7c6-423b-b6fa-811e0c27b3e0", com.azure.core.util.Context.NONE);
    }
}
