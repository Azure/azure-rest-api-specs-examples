/** Samples for WebPubSub ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSub_ListSkus.json
     */
    /**
     * Sample code: WebPubSub_ListSkus.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubListSkus(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubs()
            .listSkusWithResponse("myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
