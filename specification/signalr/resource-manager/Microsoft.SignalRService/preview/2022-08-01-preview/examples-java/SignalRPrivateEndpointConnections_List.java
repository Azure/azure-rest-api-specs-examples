import com.azure.core.util.Context;

/** Samples for SignalRPrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/SignalRPrivateEndpointConnections_List.json
     */
    /**
     * Sample code: SignalRPrivateEndpointConnections_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateEndpointConnectionsList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRPrivateEndpointConnections().list("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
