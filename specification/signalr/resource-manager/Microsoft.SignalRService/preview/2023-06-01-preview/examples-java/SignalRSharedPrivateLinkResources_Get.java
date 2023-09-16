/** Samples for SignalRSharedPrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRSharedPrivateLinkResources_Get.json
     */
    /**
     * Sample code: SignalRSharedPrivateLinkResources_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRSharedPrivateLinkResourcesGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRSharedPrivateLinkResources()
            .getWithResponse("upstream", "myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
