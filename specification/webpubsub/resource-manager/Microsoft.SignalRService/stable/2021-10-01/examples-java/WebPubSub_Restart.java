import com.azure.core.util.Context;

/** Samples for WebPubSub Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_Restart.json
     */
    /**
     * Sample code: WebPubSub_Restart.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubRestart(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().restart("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
