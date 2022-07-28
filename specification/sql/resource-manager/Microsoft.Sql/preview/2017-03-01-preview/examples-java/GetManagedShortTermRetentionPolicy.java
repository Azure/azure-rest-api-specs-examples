import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/** Samples for ManagedBackupShortTermRetentionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/GetManagedShortTermRetentionPolicy.json
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
            .getManagedBackupShortTermRetentionPolicies()
            .getWithResponse(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "testdb",
                ManagedShortTermRetentionPolicyName.DEFAULT,
                Context.NONE);
    }
}
