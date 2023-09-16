/** Samples for SignalRReplicas List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRReplicas_List.json
     */
    /**
     * Sample code: SignalRReplicas_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRReplicasList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRReplicas().list("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
