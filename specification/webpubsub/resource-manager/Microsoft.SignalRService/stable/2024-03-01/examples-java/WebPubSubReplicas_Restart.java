
/**
 * Samples for WebPubSubReplicas Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubReplicas_Restart.json
     */
    /**
     * Sample code: WebPubSubReplicas_Restart.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubReplicasRestart(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubReplicas().restart("myResourceGroup", "myWebPubSubService", "myWebPubSubService-eastus",
            com.azure.core.util.Context.NONE);
    }
}
