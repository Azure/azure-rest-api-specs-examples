/** Samples for WebPubSubReplicas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubReplicas_Get.json
     */
    /**
     * Sample code: WebPubSubReplicas_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubReplicasGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubReplicas()
            .getWithResponse(
                "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", com.azure.core.util.Context.NONE);
    }
}
