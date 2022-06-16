import com.azure.core.util.Context;

/** Samples for SignalRPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRPrivateLinkResources_List.json
     */
    /**
     * Sample code: SignalRPrivateLinkResources_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateLinkResourcesList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRPrivateLinkResources().list("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
