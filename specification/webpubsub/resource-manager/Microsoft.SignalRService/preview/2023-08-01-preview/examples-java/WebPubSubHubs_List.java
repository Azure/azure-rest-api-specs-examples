/** Samples for WebPubSubHubs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubHubs_List.json
     */
    /**
     * Sample code: WebPubSubHubs_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubHubsList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubHubs().list("myResourceGroup", "myWebPubSubService", com.azure.core.util.Context.NONE);
    }
}
