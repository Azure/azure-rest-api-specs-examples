
/**
 * Samples for WebPubSubReplicaSharedPrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubReplicaSharedPrivateLinkResources_List.json
     */
    /**
     * Sample code: WebPubSubReplicaSharedPrivateLinkResources_List.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void
        webPubSubReplicaSharedPrivateLinkResourcesList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubReplicaSharedPrivateLinkResources().list("myResourceGroup", "myWebPubSubService",
            "myWebPubSubService-eastus", com.azure.core.util.Context.NONE);
    }
}
