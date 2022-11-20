import com.azure.core.util.Context;

/** Samples for WebPubSubCustomCertificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubCustomCertificates_Get.json
     */
    /**
     * Sample code: WebPubSubCustomCertificates_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomCertificatesGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomCertificates()
            .getWithResponse("myResourceGroup", "myWebPubSubService", "myCert", Context.NONE);
    }
}
