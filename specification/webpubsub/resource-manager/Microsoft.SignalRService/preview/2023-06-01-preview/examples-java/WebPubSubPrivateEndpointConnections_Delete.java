/** Samples for WebPubSubPrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubPrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: WebPubSubPrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubPrivateEndpointConnectionsDelete(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubPrivateEndpointConnections()
            .delete(
                "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                "myResourceGroup",
                "myWebPubSubService",
                com.azure.core.util.Context.NONE);
    }
}
