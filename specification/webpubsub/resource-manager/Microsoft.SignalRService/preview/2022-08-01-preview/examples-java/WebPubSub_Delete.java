import com.azure.core.util.Context;

/** Samples for WebPubSub Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSub_Delete.json
     */
    /**
     * Sample code: WebPubSub_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubDelete(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().delete("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
