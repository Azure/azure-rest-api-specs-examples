
/**
 * Samples for WebPubSubCustomCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubCustomCertificates_Get.json
     */
    /**
     * Sample code: WebPubSubCustomCertificates_Get.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomCertificatesGet(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubCustomCertificates().getWithResponse("myResourceGroup", "myWebPubSubService", "myCert",
            com.azure.core.util.Context.NONE);
    }
}
