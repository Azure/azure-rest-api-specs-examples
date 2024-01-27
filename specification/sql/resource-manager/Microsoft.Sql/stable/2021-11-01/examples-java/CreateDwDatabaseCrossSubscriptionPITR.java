
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import java.time.OffsetDateTime;

/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateDwDatabaseCrossSubscriptionPITR
     * .json
     */
    /**
     * Sample code: Creates a data warehouse database as a cross-subscription restore from a restore point of an
     * existing database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADataWarehouseDatabaseAsACrossSubscriptionRestoreFromARestorePointOfAnExistingDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdw",
            new DatabaseInner().withLocation("southeastasia").withCreateMode(CreateMode.POINT_IN_TIME_RESTORE)
                .withRestorePointInTime(OffsetDateTime.parse("2022-01-22T05:35:31.503Z")).withSourceResourceId(
                    "/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/srcsvr/databases/srcdw"),
            Context.NONE);
    }
}
