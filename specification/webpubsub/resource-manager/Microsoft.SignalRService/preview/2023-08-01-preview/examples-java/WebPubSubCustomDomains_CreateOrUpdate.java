import com.azure.resourcemanager.webpubsub.models.ResourceReference;

/** Samples for WebPubSubCustomDomains CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubCustomDomains_CreateOrUpdate.json
     */
    /**
     * Sample code: WebPubSubCustomDomains_CreateOrUpdate.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCustomDomainsCreateOrUpdate(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubCustomDomains()
            .define("myDomain")
            .withExistingWebPubSub("myResourceGroup", "myWebPubSubService")
            .withDomainName("example.com")
            .withCustomCertificate(
                new ResourceReference()
                    .withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/customCertificates/myCert"))
            .create();
    }
}
