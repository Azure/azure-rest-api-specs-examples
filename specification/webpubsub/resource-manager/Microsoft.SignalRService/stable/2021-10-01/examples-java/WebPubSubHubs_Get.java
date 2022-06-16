import com.azure.core.util.Context;

/** Samples for WebPubSubHubs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_Get.json
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
