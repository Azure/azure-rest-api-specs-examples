import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.models.SapApplicationServerInstance;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapApplicationServerInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPApplicationServerInstances_Update.json
     */
    /**
     * Sample code: SAPApplicationServerInstances_Update.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPApplicationServerInstancesUpdate(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        SapApplicationServerInstance resource =
            manager.sapApplicationServerInstances().getWithResponse("test-rg", "X00", "app01", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1")).apply();
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
