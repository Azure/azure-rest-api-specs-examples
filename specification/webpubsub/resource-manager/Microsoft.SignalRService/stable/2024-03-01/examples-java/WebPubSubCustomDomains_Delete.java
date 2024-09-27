
/**
 * Samples for WebPubSubCustomDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubCustomDomains_Delete.json
     */
    /**
     * Sample code: WebPubSubCustomDomains_Delete.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomDomainsDelete(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubCustomDomains().delete("myResourceGroup", "myWebPubSubService", "example",
            com.azure.core.util.Context.NONE);
    }
}
