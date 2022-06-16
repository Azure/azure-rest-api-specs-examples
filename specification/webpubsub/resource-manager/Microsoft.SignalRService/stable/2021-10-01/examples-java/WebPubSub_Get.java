import com.azure.core.util.Context;

/** Samples for WebPubSub GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_Get.json
     */
    /**
     * Sample code: WebPubSub_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().getByResourceGroupWithResponse("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
