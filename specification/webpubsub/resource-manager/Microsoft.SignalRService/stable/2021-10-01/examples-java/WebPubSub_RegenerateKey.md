Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-webpubsub_1.0.0-beta.2/sdk/webpubsub/azure-resourcemanager-webpubsub/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.webpubsub.models.KeyType;
import com.azure.resourcemanager.webpubsub.models.RegenerateKeyParameters;

/** Samples for WebPubSub RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_RegenerateKey.json
     */
    /**
     * Sample code: WebPubSub_RegenerateKey.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubRegenerateKey(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubs()
            .regenerateKey(
                "myResourceGroup",
                "myWebPubSubService",
                new RegenerateKeyParameters().withKeyType(KeyType.PRIMARY),
                Context.NONE);
    }
}
```
