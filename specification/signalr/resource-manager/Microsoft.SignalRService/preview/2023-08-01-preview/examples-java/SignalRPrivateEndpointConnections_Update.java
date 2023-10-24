import com.azure.resourcemanager.signalr.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.signalr.models.PrivateEndpoint;
import com.azure.resourcemanager.signalr.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.signalr.models.PrivateLinkServiceConnectionStatus;

/** Samples for SignalRPrivateEndpointConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRPrivateEndpointConnections_Update.json
     */
    /**
     * Sample code: SignalRPrivateEndpointConnections_Update.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateEndpointConnectionsUpdate(
        com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRPrivateEndpointConnections()
            .updateWithResponse(
                "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                "myResourceGroup",
                "mySignalRService",
                new PrivateEndpointConnectionInner()
                    .withPrivateEndpoint(new PrivateEndpoint())
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                            .withActionsRequired("None")),
                com.azure.core.util.Context.NONE);
    }
}
