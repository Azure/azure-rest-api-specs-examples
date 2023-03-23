/** Samples for WebPubSubCustomDomains List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubCustomDomains_List.json
     */
    /**
     * Sample code: WebPubSubCustomDomains_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomDomainsList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomDomains()
            .list("myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
