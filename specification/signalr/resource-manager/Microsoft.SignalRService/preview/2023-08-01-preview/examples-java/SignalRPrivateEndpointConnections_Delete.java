/** Samples for SignalRPrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRPrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: SignalRPrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateEndpointConnectionsDelete(
        com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRPrivateEndpointConnections()
            .delete(
                "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                "myResourceGroup",
                "mySignalRService",
                com.azure.core.util.Context.NONE);
    }
}
