import com.azure.core.util.Context;

/** Samples for WebPubSubHubs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_List.json
     */
    /**
     * Sample code: WebPubSubHubs_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubHubsList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubHubs().list("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
