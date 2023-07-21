/** Samples for WebPubSubCustomCertificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubCustomCertificates_Get.json
     */
    /**
     * Sample code: WebPubSubCustomCertificates_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomCertificatesGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomCertificates()
            .getWithResponse("myResourceGroup", "myWebPubSubService", "myCert", com.azure.core.util.Context.NONE);
    }
}
