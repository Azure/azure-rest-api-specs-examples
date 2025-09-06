
import com.azure.resourcemanager.iotfirmwaredefense.models.Workspace;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Workspaces_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        workspacesUpdateMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        Workspace resource = manager.workspaces()
            .getByResourceGroupWithResponse("rgworkspaces", "WorkspaceName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
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
