/** Samples for SignalRReplicas Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRReplicas_Delete.json
     */
    /**
     * Sample code: SignalRReplicas_Delete.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRReplicasDelete(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRReplicas()
            .deleteWithResponse(
                "myResourceGroup", "mySignalRService", "mySignalRService-eastus", com.azure.core.util.Context.NONE);
    }
}
