import com.azure.core.util.Context;

/** Samples for WebPubSubPrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubPrivateEndpointConnections_List.json
     */
    /**
     * Sample code: WebPubSubPrivateEndpointConnections_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubPrivateEndpointConnectionsList(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubPrivateEndpointConnections().list("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
