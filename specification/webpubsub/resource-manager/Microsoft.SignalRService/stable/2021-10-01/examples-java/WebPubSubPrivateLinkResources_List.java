import com.azure.core.util.Context;

/** Samples for WebPubSubPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateLinkResources_List.json
     */
    /**
     * Sample code: WebPubSubPrivateLinkResources_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubPrivateLinkResourcesList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubPrivateLinkResources().list("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
