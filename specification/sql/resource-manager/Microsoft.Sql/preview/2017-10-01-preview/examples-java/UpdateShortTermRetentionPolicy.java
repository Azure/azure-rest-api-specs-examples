import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.BackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ShortTermRetentionPolicyName;

/** Samples for BackupShortTermRetentionPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/UpdateShortTermRetentionPolicy.json
     */
    /**
     * Sample code: Update the short term retention policy for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheShortTermRetentionPolicyForTheDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getBackupShortTermRetentionPolicies()
            .createOrUpdate(
                "resourceGroup",
                "testsvr",
                "testdb",
                ShortTermRetentionPolicyName.DEFAULT,
                new BackupShortTermRetentionPolicyInner().withRetentionDays(14),
                Context.NONE);
    }
}
