import com.azure.core.util.Context;

/** Samples for SignalR Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_Restart.json
     */
    /**
     * Sample code: SignalR_Restart.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRRestart(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().restart("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
