
import com.azure.resourcemanager.durabletask.models.Scheduler;
import com.azure.resourcemanager.durabletask.models.SchedulerPropertiesUpdate;
import com.azure.resourcemanager.durabletask.models.SchedulerSkuUpdate;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Schedulers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/Schedulers_Update.json
     */
    /**
     * Sample code: Schedulers_Update.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        Scheduler resource = manager.schedulers()
            .getByResourceGroupWithResponse("rgopenapi", "testscheduler", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("hello", "world"))
            .withProperties(new SchedulerPropertiesUpdate().withIpAllowlist(Arrays.asList("10.0.0.0/8"))
                .withSku(new SchedulerSkuUpdate().withName("Dedicated").withCapacity(3)))
            .apply();
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
