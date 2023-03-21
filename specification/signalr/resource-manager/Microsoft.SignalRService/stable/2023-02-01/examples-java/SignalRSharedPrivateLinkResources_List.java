/** Samples for SignalRSharedPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalRSharedPrivateLinkResources_List.json
     */
    /**
     * Sample code: SignalRSharedPrivateLinkResources_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRSharedPrivateLinkResourcesList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRSharedPrivateLinkResources()
            .list("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
