
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseInner;
import com.azure.resourcemanager.sql.models.ManagedDatabaseCreateMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedDatabases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseCreateRecovery.json
     */
    /**
     * Sample code: Creates a new managed database from restoring a geo-replicated backup.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createsANewManagedDatabaseFromRestoringAGeoReplicatedBackup(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().createOrUpdate("Default-SQL-SouthEastAsia", "server1",
            "testdb_recovered",
            new ManagedDatabaseInner().withLocation("southeastasia").withCreateMode(ManagedDatabaseCreateMode.RECOVERY)
                .withRecoverableDatabaseId(
                    "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/Default-SQL-WestEurope/providers/Microsoft.Sql/managedInstances/testsvr/recoverableDatabases/testdb"),
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
