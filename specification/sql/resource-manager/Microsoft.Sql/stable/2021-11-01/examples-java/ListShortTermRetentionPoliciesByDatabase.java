
import com.azure.core.util.Context;

/** Samples for BackupShortTermRetentionPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ListShortTermRetentionPoliciesByDatabase.json
     */
    /**
     * Sample code: Get the short term retention policy for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getBackupShortTermRetentionPolicies()
            .listByDatabase("Default-SQL-SouthEastAsia", "testsvr", "testdb", Context.NONE);
    }
}
