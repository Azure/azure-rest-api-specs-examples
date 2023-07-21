import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Create_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesCreateMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .workspaces()
            .define("E___-3")
            .withRegion("jjwbseilitjgdrhbvvkwviqj")
            .withExistingResourceGroup("rgworkspaces")
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
