import com.azure.core.util.Context;

/** Samples for SignalRPrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalRPrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: SignalRPrivateEndpointConnections_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateEndpointConnectionsGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRPrivateEndpointConnections()
            .getWithResponse(
                "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                "myResourceGroup",
                "mySignalRService",
                Context.NONE);
    }
}
