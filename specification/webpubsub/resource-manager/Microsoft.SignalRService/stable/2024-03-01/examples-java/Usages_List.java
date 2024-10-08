
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/Usages_List.json
     */
    /**
     * Sample code: Usages_List.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void usagesList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.usages().list("eastus", com.azure.core.util.Context.NONE);
    }
}
