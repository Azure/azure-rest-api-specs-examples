import com.azure.core.util.Context;

/** Samples for WebPubSub GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSub_Get.json
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
