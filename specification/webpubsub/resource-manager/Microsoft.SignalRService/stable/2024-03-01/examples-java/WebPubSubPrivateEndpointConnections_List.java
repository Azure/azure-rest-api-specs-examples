
/**
 * Samples for WebPubSubPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubPrivateEndpointConnections_List.json
     */
    /**
     * Sample code: WebPubSubPrivateEndpointConnections_List.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void
        webPubSubPrivateEndpointConnectionsList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubPrivateEndpointConnections().list("myResourceGroup", "myWebPubSubService",
            com.azure.core.util.Context.NONE);
    }
}
