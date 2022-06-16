import com.azure.core.util.Context;

/** Samples for WebPubSubPrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Delete.json
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
                Context.NONE);
    }
}
