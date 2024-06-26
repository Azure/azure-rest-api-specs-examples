import com.azure.core.util.Context;
import com.azure.resourcemanager.networkfunction.models.CollectorPolicy;
import java.util.HashMap;
import java.util.Map;

/** Samples for CollectorPolicies UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyUpdateTags.json
     */
    /**
     * Sample code: Update Collector Policy tags.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void updateCollectorPolicyTags(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        CollectorPolicy resource =
            manager.collectorPolicies().getWithResponse("rg1", "atc", "cp1", Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "value1", "key2", "value2")).apply();
    }

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
