import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ShortTermRetentionPolicyName;

/** Samples for BackupShortTermRetentionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/GetShortTermRetentionPolicy.json
     */
    /**
     * Sample code: Get the short term retention policy for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheShortTermRetentionPolicyForTheDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getBackupShortTermRetentionPolicies()
            .getWithResponse(
                "Default-SQL-SouthEastAsia", "testsvr", "testdb", ShortTermRetentionPolicyName.DEFAULT, Context.NONE);
    }
}
