import java.util.HashMap;
import java.util.Map;

/** Samples for Monitors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_CreateOrUpdate_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .define("myMonitor")
            .withRegion("West US 2")
            .withExistingResourceGroup("myResourceGroup")
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
