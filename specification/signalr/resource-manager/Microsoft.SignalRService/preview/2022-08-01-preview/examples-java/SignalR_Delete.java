import com.azure.core.util.Context;

/** Samples for SignalR Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/SignalR_Delete.json
     */
    /**
     * Sample code: SignalR_Delete.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRDelete(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().delete("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
