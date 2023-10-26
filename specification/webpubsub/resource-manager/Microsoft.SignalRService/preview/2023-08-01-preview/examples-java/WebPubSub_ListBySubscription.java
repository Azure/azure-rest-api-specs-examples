/** Samples for WebPubSub List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSub_ListBySubscription.json
     */
    /**
     * Sample code: WebPubSub_ListBySubscription.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubListBySubscription(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().list(com.azure.core.util.Context.NONE);
    }
}
