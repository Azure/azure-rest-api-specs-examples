/** Samples for WebPubSubCustomDomains Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubCustomDomains_Get.json
     */
    /**
     * Sample code: WebPubSubCustomDomains_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomDomainsGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomDomains()
            .getWithResponse("myResourceGroup", "myWebPubSubService", "example", com.azure.core.util.Context.NONE);
    }
}
