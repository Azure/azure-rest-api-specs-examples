import com.azure.core.util.Context;

/** Samples for SignalR ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_ListSkus.json
     */
    /**
     * Sample code: SignalR_ListSkus.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListSkus(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().listSkusWithResponse("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
