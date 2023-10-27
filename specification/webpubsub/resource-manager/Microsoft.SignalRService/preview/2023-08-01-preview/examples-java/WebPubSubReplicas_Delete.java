/** Samples for WebPubSubReplicas Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_Delete.json
     */
    /**
     * Sample code: WebPubSubReplicas_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubReplicasDelete(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubReplicas()
            .deleteWithResponse(
                "myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus", com.azure.core.util.Context.NONE);
    }
}
