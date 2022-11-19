import com.azure.core.util.Context;

/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/Usages_List.json
     */
    /**
     * Sample code: Usages_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void usagesList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.usages().list("eastus", Context.NONE);
    }
}
