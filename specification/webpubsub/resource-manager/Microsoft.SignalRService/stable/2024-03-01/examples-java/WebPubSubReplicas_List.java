
/**
 * Samples for WebPubSubReplicas List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubReplicas_List.json
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
