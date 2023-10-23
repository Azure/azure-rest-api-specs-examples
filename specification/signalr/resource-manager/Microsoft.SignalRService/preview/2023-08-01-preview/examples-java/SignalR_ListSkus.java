/** Samples for SignalR ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalR_ListSkus.json
     */
    /**
     * Sample code: SignalR_ListSkus.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListSkus(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRs()
            .listSkusWithResponse("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
