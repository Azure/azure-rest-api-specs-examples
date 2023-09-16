/** Samples for SignalR ListReplicaSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalR_ListReplicaSkus.json
     */
    /**
     * Sample code: SignalR_ListReplicaSkus.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListReplicaSkus(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRs()
            .listReplicaSkusWithResponse(
                "myResourceGroup", "mySignalRService", "mySignalRService-eastus", com.azure.core.util.Context.NONE);
    }
}
