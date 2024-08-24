
import com.azure.resourcemanager.sql.fluent.models.DatabaseAutomaticTuningInner;
import com.azure.resourcemanager.sql.models.AutomaticTuningMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DatabaseAutomaticTuning Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseAutomaticTuningUpdateMin.json
     */
    /**
     * Sample code: Updates database automatic tuning settings with minimal properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesDatabaseAutomaticTuningSettingsWithMinimalProperties(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAutomaticTunings().updateWithResponse(
            "default-sql-onebox", "testsvr11", "db1",
            new DatabaseAutomaticTuningInner().withDesiredState(AutomaticTuningMode.AUTO),
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
