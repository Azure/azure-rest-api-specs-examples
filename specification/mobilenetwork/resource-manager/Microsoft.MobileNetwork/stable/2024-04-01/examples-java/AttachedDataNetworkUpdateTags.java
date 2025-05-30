
import com.azure.resourcemanager.mobilenetwork.models.AttachedDataNetwork;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AttachedDataNetworks UpdateTags.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * AttachedDataNetworkUpdateTags.json
     */
    /**
     * Sample code: Update attached data network tags.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        updateAttachedDataNetworkTags(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        AttachedDataNetwork resource = manager.attachedDataNetworks().getWithResponse("rg1", "TestPacketCoreCP",
            "TestPacketCoreDP", "TestAttachedDataNetwork", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
