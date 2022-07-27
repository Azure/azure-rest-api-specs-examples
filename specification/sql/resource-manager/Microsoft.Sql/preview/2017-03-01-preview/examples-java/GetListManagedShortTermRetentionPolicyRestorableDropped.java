import com.azure.core.util.Context;

/** Samples for ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies ListByRestorableDroppedDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/GetListManagedShortTermRetentionPolicyRestorableDropped.json
     */
    /**
     * Sample code: Get the short term retention policy list for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheShortTermRetentionPolicyListForTheDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies()
            .listByRestorableDroppedDatabase(
                "Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000", Context.NONE);
    }
}
