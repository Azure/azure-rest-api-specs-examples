/** Samples for SignalR List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalR_ListBySubscription.json
     */
    /**
     * Sample code: SignalR_ListBySubscription.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListBySubscription(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().list(com.azure.core.util.Context.NONE);
    }
}
