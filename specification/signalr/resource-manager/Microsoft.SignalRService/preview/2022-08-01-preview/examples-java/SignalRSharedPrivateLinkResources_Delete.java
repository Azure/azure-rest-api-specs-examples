import com.azure.core.util.Context;

/** Samples for SignalRSharedPrivateLinkResources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/SignalRSharedPrivateLinkResources_Delete.json
     */
    /**
     * Sample code: SignalRSharedPrivateLinkResources_Delete.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRSharedPrivateLinkResourcesDelete(
        com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRSharedPrivateLinkResources()
            .delete("upstream", "myResourceGroup", "mySignalRService", Context.NONE);
    }
}
