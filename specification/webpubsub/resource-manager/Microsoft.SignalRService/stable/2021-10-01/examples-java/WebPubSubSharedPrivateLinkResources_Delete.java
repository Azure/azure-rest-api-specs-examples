import com.azure.core.util.Context;

/** Samples for WebPubSubSharedPrivateLinkResources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubSharedPrivateLinkResources_Delete.json
     */
    /**
     * Sample code: WebPubSubSharedPrivateLinkResources_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubSharedPrivateLinkResourcesDelete(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubSharedPrivateLinkResources()
            .delete("upstream", "myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
