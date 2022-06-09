```java
import com.azure.core.util.Context;

/** Samples for WebPubSub ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_ListByResourceGroup.json
     */
    /**
     * Sample code: WebPubSub_ListByResourceGroup.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubListByResourceGroup(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubs().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-webpubsub_1.0.0-beta.2/sdk/webpubsub/azure-resourcemanager-webpubsub/README.md) on how to add the SDK to your project and authenticate.
