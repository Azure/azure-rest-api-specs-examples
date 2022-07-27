import com.azure.core.util.Context;

/** Samples for BackupLongTermRetentionPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/LongTermRetentionPolicyGet.json
     */
    /**
     * Sample code: Get the long term retention policy for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheLongTermRetentionPolicyForTheDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getBackupLongTermRetentionPolicies()
            .listByDatabaseWithResponse("resourceGroup", "testserver", "testDatabase", Context.NONE);
    }
}
