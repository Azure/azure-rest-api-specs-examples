
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * CreateDwDatabaseCrossSubscriptionRestore.json
     */
    /**
     * Sample code: Creates a data warehouse database as a cross-subscription restore from a backup of a dropped
     * database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsADataWarehouseDatabaseAsACrossSubscriptionRestoreFromABackupOfADroppedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
            "testsvr", "testdw",
            new DatabaseInner().withLocation("southeastasia").withCreateMode(CreateMode.RESTORE).withSourceResourceId(
                "/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/srcsvr/restorableDroppedDatabases/srcdw,131403269876900000"),
            com.azure.core.util.Context.NONE);
    }
}
