/** Samples for WebPubSubSharedPrivateLinkResources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubSharedPrivateLinkResources_Delete.json
     */
    /**
     * Sample code: WebPubSubSharedPrivateLinkResources_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubSharedPrivateLinkResourcesDelete(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubSharedPrivateLinkResources()
            .delete("upstream", "myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
