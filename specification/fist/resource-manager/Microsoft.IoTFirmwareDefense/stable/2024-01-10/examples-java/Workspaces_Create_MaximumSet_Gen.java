
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Workspaces_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        workspacesCreateMaximumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().define("E___-3").withRegion("jjwbseilitjgdrhbvvkwviqj")
            .withExistingResourceGroup("rgworkspaces").withTags(mapOf("key450", "fakeTokenPlaceholder")).create();
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
