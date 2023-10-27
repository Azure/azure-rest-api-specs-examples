/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void operationsList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
