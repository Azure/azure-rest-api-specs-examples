/** Samples for WebPubSubReplicas List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubReplicas_List.json
     */
    /**
     * Sample code: WebPubSubReplicas_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubReplicasList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubReplicas().list("myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
