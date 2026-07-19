
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseInner;
import com.azure.resourcemanager.sql.models.ManagedDatabaseCreateMode;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedDatabases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseCreateCrossSubscriptionPointInTimeRestore.json
     */
    /**
     * Sample code: Creates a new managed database using cross subscription point in time restore.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createsANewManagedDatabaseUsingCrossSubscriptionPointInTimeRestore(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().createOrUpdate("Default-SQL-SouthEastAsia", "managedInstance",
            "managedDatabase",
            new ManagedDatabaseInner().withLocation("southeastasia")
                .withRestorePointInTime(OffsetDateTime.parse("2017-07-14T05:35:31.503Z"))
                .withCreateMode(ManagedDatabaseCreateMode.POINT_IN_TIME_RESTORE)
                .withCrossSubscriptionSourceDatabaseId(
                    "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr2/databases/testdb")
                .withCrossSubscriptionTargetManagedInstanceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr"),
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
