/** Samples for SignalRPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRPrivateLinkResources_List.json
     */
    /**
     * Sample code: SignalRPrivateLinkResources_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRPrivateLinkResourcesList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRPrivateLinkResources()
            .list("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
