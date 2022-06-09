```java
import com.azure.core.util.Context;

/** Samples for WebPubSubSharedPrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubSharedPrivateLinkResources_Get.json
     */
    /**
     * Sample code: WebPubSubSharedPrivateLinkResources_Get.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubSharedPrivateLinkResourcesGet(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubSharedPrivateLinkResources()
            .getWithResponse("upstream", "myResourceGroup", "myWebPubSubService", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-webpubsub_1.0.0-beta.2/sdk/webpubsub/azure-resourcemanager-webpubsub/README.md) on how to add the SDK to your project and authenticate.
