import com.azure.core.util.Context;

/** Samples for WebPubSubHubs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubHubs_Get.json
     */
    /**
     * Sample code: WebPubSubHubs_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubHubsGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubHubs().getWithResponse("exampleHub", "myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
