
/**
 * Samples for SignalR ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/
     * SignalR_ListKeys.json
     */
    /**
     * Sample code: SignalR_ListKeys.
     * 
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListKeys(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().listKeysWithResponse("myResourceGroup", "mySignalRService",
            com.azure.core.util.Context.NONE);
    }
}
