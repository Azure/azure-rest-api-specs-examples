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
