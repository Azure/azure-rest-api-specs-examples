
/**
 * Samples for WebPubSub ListReplicaSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSub_ListReplicaSkus.json
     */
    /**
     * Sample code: WebPubSub_ListReplicaSkus.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubListReplicaSkus(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().listReplicaSkusWithResponse("myResourceGroup", "myWebPubSubService",
            "myWebPubSubService-eastus", com.azure.core.util.Context.NONE);
    }
}
