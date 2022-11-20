import com.azure.core.util.Context;

/** Samples for WebPubSubSharedPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubSharedPrivateLinkResources_List.json
     */
    /**
     * Sample code: WebPubSubSharedPrivateLinkResources_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubSharedPrivateLinkResourcesList(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubSharedPrivateLinkResources().list("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
