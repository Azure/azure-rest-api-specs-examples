import java.util.HashMap;
import java.util.Map;

/** Samples for SapApplicationServerInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_Create_HA_AvSet.json
     */
    /**
     * Sample code: Create SAP Application Server Instances for HA System with Availability Set.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createSAPApplicationServerInstancesForHASystemWithAvailabilitySet(
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
