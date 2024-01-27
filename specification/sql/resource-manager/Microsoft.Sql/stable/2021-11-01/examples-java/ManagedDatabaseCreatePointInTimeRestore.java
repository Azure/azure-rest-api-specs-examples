
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseInner;
import com.azure.resourcemanager.sql.models.ManagedDatabaseCreateMode;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedDatabases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseCreatePointInTimeRestore.json
     */
    /**
     * Sample code: Creates a new managed database using point in time restore.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsANewManagedDatabaseUsingPointInTimeRestore(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "managedInstance", "managedDatabase",
            new ManagedDatabaseInner().withLocation("southeastasia")
                .withRestorePointInTime(OffsetDateTime.parse("2017-07-14T05:35:31.503Z"))
                .withCreateMode(ManagedDatabaseCreateMode.POINT_IN_TIME_RESTORE).withSourceDatabaseId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/databases/testdb"),
            Context.NONE);
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
