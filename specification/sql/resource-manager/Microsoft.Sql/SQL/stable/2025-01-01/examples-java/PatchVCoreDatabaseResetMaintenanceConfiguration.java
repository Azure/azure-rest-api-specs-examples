
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PatchVCoreDatabaseResetMaintenanceConfiguration.json
     */
    /**
     * Sample code: Resets maintenance window of a database to default.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        resetsMaintenanceWindowOfADatabaseToDefault(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().update("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new DatabaseUpdate().withSku(new Sku().withName("BC_Gen5_4")).withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
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
