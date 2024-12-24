
/**
 * Samples for SignalRPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/
     * SignalRPrivateEndpointConnections_List.json
     */
    /**
     * Sample code: SignalRPrivateEndpointConnections_List.
     * 
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateEndpointConnectionsList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRPrivateEndpointConnections().list("myResourceGroup", "mySignalRService",
            com.azure.core.util.Context.NONE);
    }
}
