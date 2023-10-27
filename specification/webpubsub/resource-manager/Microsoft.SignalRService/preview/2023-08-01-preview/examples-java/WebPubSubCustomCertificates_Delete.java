/** Samples for WebPubSubCustomCertificates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubCustomCertificates_Delete.json
     */
    /**
     * Sample code: WebPubSubCustomCertificates_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomCertificatesDelete(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomCertificates()
            .deleteWithResponse("myResourceGroup", "myWebPubSubService", "myCert", com.azure.core.util.Context.NONE);
    }
}
