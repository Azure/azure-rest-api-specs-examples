import java.util.HashMap;
import java.util.Map;

/** Samples for SapApplicationServerInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPApplicationServerInstances_Create.json
     */
    /**
     * Sample code: SAPApplicationServerInstances_Create.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPApplicationServerInstancesCreate(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapApplicationServerInstances()
            .define("app01")
            .withRegion("westcentralus")
            .withExistingSapVirtualInstance("test-rg", "X00")
            .withTags(mapOf())
            .create();
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
