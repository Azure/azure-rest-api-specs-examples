
import com.azure.resourcemanager.sql.fluent.models.ServerAutomaticTuningInner;
import com.azure.resourcemanager.sql.models.AutomaticTuningServerMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ServerAutomaticTuning Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAutomaticTuningUpdateMin.json
     */
    /**
     * Sample code: Updates server automatic tuning settings with minimal properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updatesServerAutomaticTuningSettingsWithMinimalProperties(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAutomaticTunings().updateWithResponse("default-sql-onebox", "testsvr11",
            new ServerAutomaticTuningInner().withDesiredState(AutomaticTuningServerMode.AUTO),
            com.azure.core.util.Context.NONE);
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
