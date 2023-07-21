/** Samples for WebPubSubPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubPrivateLinkResources_List.json
     */
    /**
     * Sample code: WebPubSubPrivateLinkResources_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubPrivateLinkResourcesList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubPrivateLinkResources()
            .list("myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
