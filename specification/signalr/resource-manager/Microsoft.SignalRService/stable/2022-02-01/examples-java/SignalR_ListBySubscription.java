import com.azure.core.util.Context;

/** Samples for SignalR List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_ListBySubscription.json
     */
    /**
     * Sample code: SignalR_ListBySubscription.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListBySubscription(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().list(Context.NONE);
    }
}
