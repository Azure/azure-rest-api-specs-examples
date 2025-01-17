
import com.azure.resourcemanager.signalr.models.ResourceSku;
import com.azure.resourcemanager.signalr.models.SignalRSkuTier;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SignalRReplicas CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/
     * SignalRReplicas_CreateOrUpdate.json
     */
    /**
     * Sample code: SignalRReplicas_CreateOrUpdate.
     * 
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRReplicasCreateOrUpdate(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRReplicas().define("mySignalRService-eastus").withRegion("eastus")
            .withExistingSignalR("myResourceGroup", "mySignalRService").withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withSku(new ResourceSku().withName("Premium_P1").withTier(SignalRSkuTier.PREMIUM).withCapacity(1))
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
