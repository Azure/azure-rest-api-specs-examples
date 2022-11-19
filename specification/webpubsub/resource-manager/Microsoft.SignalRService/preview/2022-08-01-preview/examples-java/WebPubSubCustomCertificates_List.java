import com.azure.core.util.Context;

/** Samples for WebPubSubCustomCertificates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubCustomCertificates_List.json
     */
    /**
     * Sample code: WebPubSubCustomCertificates_List.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomCertificatesList(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubCustomCertificates().list("myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
