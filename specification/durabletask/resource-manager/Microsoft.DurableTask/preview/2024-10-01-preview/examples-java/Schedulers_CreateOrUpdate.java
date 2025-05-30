
import com.azure.resourcemanager.durabletask.models.SchedulerProperties;
import com.azure.resourcemanager.durabletask.models.SchedulerSku;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Schedulers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/Schedulers_CreateOrUpdate.json
     */
    /**
     * Sample code: Schedulers_CreateOrUpdate.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersCreateOrUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().define("testscheduler").withRegion("northcentralus").withExistingResourceGroup("rgopenapi")
            .withTags(mapOf("key7131", "fakeTokenPlaceholder", "key2138", "fakeTokenPlaceholder"))
            .withProperties(new SchedulerProperties().withIpAllowlist(Arrays.asList("10.0.0.0/8"))
                .withSku(new SchedulerSku().withName("Dedicated")))
            .create();
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
