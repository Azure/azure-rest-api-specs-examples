import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.MonitorResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for Monitors Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsUpdateMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        MonitorResource resource =
            manager.monitors().getByResourceGroupWithResponse("myResourceGroup", "myMonitor", Context.NONE).getValue();
        resource.update().apply();
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
