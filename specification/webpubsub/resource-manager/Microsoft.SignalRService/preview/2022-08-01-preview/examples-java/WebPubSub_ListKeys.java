import com.azure.core.util.Context;

/** Samples for WebPubSub ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSub_ListKeys.json
     */
    /**
     * Sample code: WebPubSub_ListKeys.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubListKeys(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().listKeysWithResponse("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
