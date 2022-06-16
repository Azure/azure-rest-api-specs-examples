import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void operationsList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.operations().list(Context.NONE);
    }
}
