
import com.azure.resourcemanager.webpubsub.models.ResourceSku;
import com.azure.resourcemanager.webpubsub.models.WebPubSubSkuTier;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebPubSubReplicas CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubReplicas_CreateOrUpdate.json
     */
    /**
     * Sample code: WebPubSubReplicas_CreateOrUpdate.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubReplicasCreateOrUpdate(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager.webPubSubReplicas().define("myWebPubSubService-eastus").withRegion("eastus")
            .withExistingWebPubSub("myResourceGroup", "myWebPubSubService")
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withSku(new ResourceSku().withName("Premium_P1").withTier(WebPubSubSkuTier.PREMIUM).withCapacity(1))
            .withResourceStopped("false").create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
