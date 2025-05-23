
import com.azure.resourcemanager.signalr.models.Replica;
import com.azure.resourcemanager.signalr.models.ResourceSku;
import com.azure.resourcemanager.signalr.models.SignalRSkuTier;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SignalRReplicas Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/
     * SignalRReplicas_Update.json
     */
    /**
     * Sample code: SignalRReplicas_Update.
     * 
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRReplicasUpdate(com.azure.resourcemanager.signalr.SignalRManager manager) {
        Replica resource = manager.signalRReplicas().getWithResponse("myResourceGroup", "mySignalRService",
            "mySignalRService-eastus", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withSku(new ResourceSku().withName("Premium_P1").withTier(SignalRSkuTier.PREMIUM).withCapacity(1))
            .withResourceStopped("false").apply();
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
