/** Samples for WebPubSubHubs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubHubs_Delete.json
     */
    /**
     * Sample code: WebPubSubHubs_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubHubsDelete(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubHubs()
            .delete("exampleHub", "myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
