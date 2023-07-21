/** Samples for WebPubSub ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSub_ListByResourceGroup.json
     */
    /**
     * Sample code: WebPubSub_ListByResourceGroup.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubListByResourceGroup(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
