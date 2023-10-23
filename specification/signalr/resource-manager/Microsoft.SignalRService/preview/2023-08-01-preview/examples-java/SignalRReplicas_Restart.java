/** Samples for SignalRReplicas Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRReplicas_Restart.json
     */
    /**
     * Sample code: SignalRReplicas_Restart.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRReplicasRestart(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRReplicas()
            .restart(
                "myResourceGroup", "mySignalRService", "mySignalRService-eastus", com.azure.core.util.Context.NONE);
    }
}
