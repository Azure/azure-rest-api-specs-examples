
import com.azure.core.util.Context;

/** Samples for ManagedDatabases Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseDelete.json
     */
    /**
     * Sample code: Delete managed database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabases().delete("Default-SQL-SouthEastAsia",
            "managedInstance", "testdb", Context.NONE);
    }
}
