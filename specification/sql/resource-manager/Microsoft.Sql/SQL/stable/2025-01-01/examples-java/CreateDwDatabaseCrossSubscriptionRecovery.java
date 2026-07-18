
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateDwDatabaseCrossSubscriptionRecovery.json
     */
    /**
     * Sample code: Creates a data warehouse database as a cross-subscription restore from a geo-backup.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createsADataWarehouseDatabaseAsACrossSubscriptionRestoreFromAGeoBackup(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().createOrUpdate("Default-SQL-WestUS", "testsvr", "testdw",
            new DatabaseInner().withLocation("westus").withCreateMode(CreateMode.RECOVERY).withSourceResourceId(
                "/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-EastUS/providers/Microsoft.Sql/servers/srcsvr/recoverabledatabases/srcdw"),
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
