import com.azure.core.util.Context;

/** Samples for SignalR ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_ListKeys.json
     */
    /**
     * Sample code: SignalR_ListKeys.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListKeys(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().listKeysWithResponse("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
