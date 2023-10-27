/** Samples for WebPubSubCustomCertificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubCustomCertificates_CreateOrUpdate.json
     */
    /**
     * Sample code: WebPubSubCustomCertificates_CreateOrUpdate.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomCertificatesCreateOrUpdate(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomCertificates()
            .define("myCert")
            .withExistingWebPubSub("myResourceGroup", "myWebPubSubService")
            .withKeyVaultBaseUri("https://myvault.keyvault.azure.net/")
            .withKeyVaultSecretName("mycert")
            .withKeyVaultSecretVersion("bb6a44b2743f47f68dad0d6cc9756432")
            .create();
    }
}
