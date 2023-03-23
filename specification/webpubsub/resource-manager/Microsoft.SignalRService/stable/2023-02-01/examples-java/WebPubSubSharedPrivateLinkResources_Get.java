/** Samples for WebPubSubSharedPrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubSharedPrivateLinkResources_Get.json
     */
    /**
     * Sample code: WebPubSubSharedPrivateLinkResources_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubSharedPrivateLinkResourcesGet(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubSharedPrivateLinkResources()
            .getWithResponse("upstream", "myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
